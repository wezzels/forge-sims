package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
)

const (
	G       = 6.67430e-11
	M_earth = 5.972e24
	R_earth = 6378137.0
	g0      = 9.80665
)

type LaunchVehicle struct {
	Name        string
	Stages      []Stage
	PayloadMass float64
}

type Stage struct {
	Thrust      float64 // Thrust in Newtons
	Isp         float64 // Specific impulse at current altitude (seconds)
	IspVac      float64 // Vacuum specific impulse (seconds)
	DryMass     float64 // Dry mass (kg)
	PropMass    float64 // Propellant mass (kg)
	FairingMass float64 // Fairing mass (kg)
}

type VehicleState struct {
	Time      float64 `json:"time"`
	Altitude  float64 `json:"altitude"`
	Downrange float64 `json:"downrange"`
	Velocity  float64 `json:"velocity"`
	VX        float64 `json:"vx"`
	VY        float64 `json:"vy"`
	Mass      float64 `json:"mass"`
	Thrust    float64 `json:"thrust"`
	Accel     float64 `json:"accel"`
	Pitch     float64 `json:"pitch"`
	StageNum  int     `json:"stage"`
	Phase     string  `json:"phase"`
	Q         float64 `json:"q"`
}

type TrajectoryResult struct {
	Simulator  string              `json:"simulator"`
	Status     string              `json:"status"`
	FinalState VehicleState       `json:"final_state"`
	Orbit      OrbitState          `json:"orbit"`
	Parameters map[string]float64 `json:"parameters"`
	Stages     []StageEvent       `json:"stage_events"`
	Milestones []Milestone         `json:"milestones"`
}

type OrbitState struct {
	Apogee       float64 `json:"apogee"`
	Perigee      float64 `json:"perigee"`
	Inclination  float64 `json:"inclination"`
	Period       float64 `json:"period"`
	Eccentricity float64 `json:"eccentricity"`
	SmAxis       float64 `json:"semi_major_axis"`
}

type StageEvent struct {
	Time       float64 `json:"time"`
	Stage      int     `json:"stage"`
	Event      string  `json:"event"`
	Altitude   float64 `json:"altitude_m"`
	Velocity   float64 `json:"velocity_ms"`
	MassBefore float64 `json:"mass_before_kg"`
	MassAfter  float64 `json:"mass_after_kg"`
}

type Milestone struct {
	Time float64 `json:"time_s"`
	Name string  `json:"name"`
	Alt  float64 `json:"altitude_km"`
	Vel  float64 `json:"velocity_ms"`
}

// GetIsp returns effective Isp based on altitude (blend between sea level and vacuum)
func (s Stage) GetIsp(altitude float64) float64 {
	// Linear blend: vacuum Isp at high altitude, sea level Isp at sea level
	if s.IspVac == 0 {
		return s.Isp // No vacuum variant
	}
	blend := math.Min(1.0, altitude/50000.0) // Full vacuum at 50km
	return s.Isp*(1-blend) + s.IspVac*blend
}

func DefaultFalcon9() LaunchVehicle {
	return LaunchVehicle{
		Name: "Falcon 9-like",
		Stages: []Stage{
			// Stage 1: Sea level Isp 282s, vacuum 311s
			{Thrust: 7.607e6, Isp: 282, IspVac: 311, DryMass: 22000, PropMass: 395700, FairingMass: 0},
			// Stage 2: Vacuum only (348s)
			{Thrust: 934000, Isp: 348, IspVac: 348, DryMass: 3900, PropMass: 107500, FairingMass: 1850},
		},
		PayloadMass: 22800,
	}
}

func DefaultAtlasV() LaunchVehicle {
	return LaunchVehicle{
		Name: "Atlas V 401",
		Stages: []Stage{
			// Stage 1: RD-180, sea level 311s, vacuum 331s
			{Thrust: 4.15e6, Isp: 280, IspVac: 331, DryMass: 20000, PropMass: 291500, FairingMass: 0},
			// Stage 2: RL10, vacuum 450s
			{Thrust: 93400, Isp: 450, IspVac: 450, DryMass: 2100, PropMass: 20800, FairingMass: 0},
		},
		PayloadMass: 5500,
	}
}

func DefaultDeltaIV() LaunchVehicle {
	return LaunchVehicle{
		Name: "Delta IV Heavy",
		Stages: []Stage{
			// Stage 1: RS-68A, sea level 360s, vacuum 414s
			{Thrust: 3.14e6, Isp: 300, IspVac: 414, DryMass: 28000, PropMass: 625000, FairingMass: 0},
			// Stage 2: RL10B-2, vacuum 462s
			{Thrust: 110000, Isp: 462, IspVac: 462, DryMass: 4100, PropMass: 19100, FairingMass: 0},
		},
		PayloadMass: 12900,
	}
}

func GravityAtAltitude(h float64) float64 {
	r := R_earth + h
	return G * M_earth / (r * r)
}

func AtmosphericDensity(h float64) float64 {
	if h > 84000 {
		return 0
	}
	if h < 11000 {
		T := 288.15 - 0.0065*h
		P := 101325.0 * math.Pow(T/288.15, 5.256)
		return P / (287.05 * T)
	}
	if h < 25000 {
		T := 216.65
		P := 22632.0 * math.Exp(-0.00009*(h-11000))
		return P / (287.05 * T)
	}
	T := 216.65 + 0.001*(h-25000)
	P := 2488.0 * math.Exp(-0.00014*(h-25000))
	return P / (287.05 * T)
}

func SimulateAscent(lv LaunchVehicle, targetAlt, targetInc, duration float64) TrajectoryResult {
	result := TrajectoryResult{
		Simulator:  "launch-veh-sim",
		Status:     "in-flight",
		Parameters: make(map[string]float64),
		Stages:     []StageEvent{},
		Milestones: []Milestone{},
	}

	// State
	state := VehicleState{
		Time:     0,
		Altitude: 0,
		VX:       0,
		VY:       0,
		Mass:     lv.PayloadMass + lv.Stages[0].DryMass + lv.Stages[0].PropMass + lv.Stages[1].DryMass + lv.Stages[1].PropMass + lv.Stages[1].FairingMass,
		StageNum: 1,
		Phase:    "liftoff",
	}

	// Track current stage (0-indexed)
	stageIdx := 0
	stagePropRemaining := []float64{lv.Stages[0].PropMass, lv.Stages[1].PropMass}

	// Physics parameters
	cd := 0.5
	area := 3.5
	pitchProgram := 0.0
	maxQ := 0.0
	maxQTime := 0.0
	tof := 0.0
	dt := 0.05

	mu := G * M_earth

	addMilestone := func(name string, t, alt, vel float64) {
		result.Milestones = append(result.Milestones, Milestone{Time: t, Name: name, Alt: alt / 1000.0, Vel: vel})
	}

	addStageEvent := func(name string, t, alt, vel, massBefore, massAfter float64) {
		result.Stages = append(result.Stages, StageEvent{
			Time: t, Stage: stageIdx + 1, Event: name,
			Altitude: alt, Velocity: vel, MassBefore: massBefore, MassAfter: massAfter,
		})
	}

	addMilestone("liftoff", 0, 0, 0)

	for tof < duration {
		stage := lv.Stages[stageIdx]

		// Check for stage 1 separation
		if stageIdx == 0 && stagePropRemaining[0] <= 0 {
			massBefore := state.Mass
			state.Mass -= stage.DryMass
			stageIdx = 1
			state.StageNum = 2
			state.Phase = "stage_separation"
			addStageEvent("stage_separation", tof, state.Altitude, state.Velocity, massBefore, state.Mass)
			addMilestone("stage_separation", tof, state.Altitude, state.Velocity)
			continue
		}

		// Get current Isp based on altitude
		isp := stage.GetIsp(state.Altitude)

		// Determine thrust
		var thrust float64
		if stagePropRemaining[stageIdx] > 0 {
			thrust = stage.Thrust
			state.Phase = "powered"
		} else {
			thrust = 0
			state.Phase = "coasting"
		}

		// Calculate current mass
		currentMass := lv.PayloadMass + stagePropRemaining[0] + lv.Stages[0].DryMass + stagePropRemaining[1] + lv.Stages[1].DryMass + lv.Stages[1].FairingMass
		if stageIdx > 0 {
			currentMass += lv.Stages[0].DryMass
		}

		// Gravity
		g := GravityAtAltitude(state.Altitude)

		// Velocity magnitude
		vTotal := math.Sqrt(state.VX*state.VX + state.VY*state.VY)

		// Dynamic pressure
		rho := AtmosphericDensity(state.Altitude)
		q := 0.5 * rho * vTotal * vTotal
		if q > maxQ {
			maxQ = q
			maxQTime = tof
		}
		state.Q = q

		// Gravity turn after max-Q
		if q < 35000 && state.Altitude > 12000 {
			if vTotal > 10 {
				desiredPitch := math.Atan2(state.VX, state.VY) * 180 / math.Pi
				pitchProgram = pitchProgram*0.97 + desiredPitch*0.03
			}
		}
		pitchRad := pitchProgram * math.Pi / 180.0

		// Drag
		drag := 0.5 * rho * vTotal * vTotal * cd * area
		dragX := 0.0
		dragY := 0.0
		if vTotal > 1 {
			dragX = drag * state.VX / vTotal / currentMass
			dragY = drag * state.VY / vTotal / currentMass
		}

		// Thrust components
		thrustX := thrust * math.Sin(pitchRad) / currentMass
		thrustY := thrust * math.Cos(pitchRad) / currentMass

		// Net acceleration
		ay := thrustY - dragY - g

		// Integrate
		state.VX += thrustX * dt - dragX*dt
		state.VY += ay * dt
		state.Altitude += state.VY * dt
		state.Downrange += state.VX * dt
		state.Time = tof
		state.Mass = currentMass
		state.Thrust = thrust
		state.Accel = math.Sqrt((thrustX-dragX)*(thrustX-dragX) + ay*ay)
		state.Pitch = pitchProgram
		state.Velocity = vTotal

		// Consume propellant
		if stagePropRemaining[stageIdx] > 0 && thrust > 0 {
			massFlowRate := thrust / (isp * g0)
			stagePropRemaining[stageIdx] -= massFlowRate * dt
			if stagePropRemaining[stageIdx] < 0 {
				stagePropRemaining[stageIdx] = 0
			}
		}

		// Check orbit achieved
		if state.Altitude >= targetAlt && state.VY > -200 && state.VY < 200 {
			r := R_earth + state.Altitude
			v := vTotal
			a := 1.0 / (2.0/r - v*v/mu)
			if a <= 0 || math.IsNaN(a) {
				result.Status = "incomplete"
				result.FinalState = state
				break
			}
			ecc := 0.0
			if a > r {
				ecc = math.Sqrt(1 - r*r/(a*a))
				if ecc > 1 {
					ecc = 1 - 1e-10
				}
			}
			apogee := a * (1 + ecc) / 1000.0
			perigee := math.Max(0, a*(1-ecc)/1000.0)
			period := 2 * math.Pi * math.Sqrt(a*a*a/mu) / 60.0

			result.Orbit = OrbitState{
				Apogee:      apogee,
				Perigee:     perigee,
				Inclination: targetInc,
				Period:      period,
				Eccentricity: ecc,
				SmAxis:      a / 1000.0,
			}
			result.Status = "orbit_achieved"
			result.FinalState = state
			addMilestone("orbit_insertion", tof, state.Altitude, vTotal)
			break
		}

		// Check crash
		if state.Altitude < 0 && state.VY < 0 {
			state.Altitude = 0
			result.Status = "crashed"
			result.FinalState = state
			addMilestone("crash", tof, 0, vTotal)
			break
		}

		tof += dt
	}

	if result.Status == "in-flight" {
		result.Status = "incomplete"
		result.FinalState = state
	}

	result.Parameters["payload_mass_kg"] = lv.PayloadMass
	result.Parameters["total_propellant_kg"] = lv.Stages[0].PropMass + lv.Stages[1].PropMass
	result.Parameters["max_q_Pa"] = maxQ
	result.Parameters["max_q_time_s"] = maxQTime
	result.Parameters["flight_time_s"] = tof

	return result
}

func main() {
	jsonFlag := flag.Bool("json", false, "Output JSON format")
	duration := flag.Float64("duration", 800, "Simulation duration (seconds)")
	targetAlt := flag.Float64("target-alt", 400000, "Target orbital altitude (meters)")
	targetInc := flag.Float64("inclination", 28.5, "Target orbit inclination (degrees)")
	vehicle := flag.String("vehicle", "falcon9", "Vehicle type (falcon9, atlas, delta)")
	flag.Parse()

	var lv LaunchVehicle
	switch *vehicle {
	case "atlas":
		lv = DefaultAtlasV()
	case "delta":
		lv = DefaultDeltaIV()
	default:
		lv = DefaultFalcon9()
	}

	result := SimulateAscent(lv, *targetAlt, *targetInc, *duration)

	if *jsonFlag {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(result)
	} else {
		fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║           LAUNCH VEHICLE SIMULATION RESULTS                     ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
		fmt.Printf("Vehicle: %s | Payload: %.0f kg\n", lv.Name, lv.PayloadMass)
		fmt.Printf("Status: %s\n\n", result.Status)

		if result.Status == "orbit_achieved" {
			fmt.Println("✓ Orbit Achieved!")
			fmt.Printf("  Apogee:    %.0f km\n", result.Orbit.Apogee)
			fmt.Printf("  Perigee:   %.0f km\n", result.Orbit.Perigee)
			fmt.Printf("  Period:    %.1f min\n", result.Orbit.Period)
			fmt.Printf("  Eccentricity: %.4f\n", result.Orbit.Eccentricity)
		}

		fmt.Printf("\nFlight Time: %.1f s\n", result.FinalState.Time)
		fmt.Printf("Final Altitude: %.0f km\n", result.FinalState.Altitude/1000)
		fmt.Printf("Final Velocity: %.0f m/s\n", result.FinalState.Velocity)
		fmt.Printf("Max Q: %.0f Pa at %.0f s\n", result.Parameters["max_q_Pa"], result.Parameters["max_q_time_s"])

		if len(result.Stages) > 0 {
			fmt.Println("\nStage Events:")
			for _, s := range result.Stages {
				fmt.Printf("  [%.1f s] Stage %d: %s (%.0f km, %.0f m/s)\n", s.Time, s.Stage, s.Event, s.Altitude/1000, s.Velocity)
			}
		}

		fmt.Println("\nMilestones:")
		for _, m := range result.Milestones {
			fmt.Printf("  [%.1f s] %s\n", m.Time, m.Name)
		}
	}
}