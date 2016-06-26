package mora

import (
	"math/rand"
	"time"
)

// JitterSigma is the standard deviation of the jitter on the interval, which
// prevents collisions by providing noise to when each node starts pinging.
const JitterSigma = 2

// MinimumWait specifies the minimum number of seconds to wait between actions.
const MinimumWait = 1 * time.Second

// JitteredInterval runs an action at an interval with some added jitter (e.g.
// plus or minus a couple of seconds) to prevent collisions with other workers.
type JitteredInterval struct {
	Action          func() error  // The action to be called on the jittered interval
	Stopped         bool          // A flag determining the state of the worker
	ShutdownChannel chan string   // A channel to communicate to the routine
	Interval        time.Duration // The interval with which to run the generator
	period          time.Duration // The actual period of the wait
}

// NewJitteredInterval creates a new jittered interval worker with the specified
// action to be called at the specified interval with some jitter.
func NewJitteredInterval(action func() error, interval time.Duration) *JitteredInterval {
	return &JitteredInterval{
		Action:          action,
		Stopped:         false,
		ShutdownChannel: make(chan string),
		Interval:        interval,
		period:          interval,
	}
}

// Run starts the interval and listens for a shutdown call.
func (jw *JitteredInterval) Run() error {

	// Loop that runs forever
	for {
		select {
		case <-jw.ShutdownChannel:
			jw.ShutdownChannel <- "Down"
			return nil
		case <-time.After(jw.period):
			// This breaks out of the select, not the for loop.
			break
		}

		// Execute the ping generator action tracking how much time went by.
		started := time.Now()
		if err := jw.Action(); err != nil {
			// If there was an error, then we're going to stop and return it.
			return err
		}

		// Reset the period less the amount of work gone by with some jitter.
		jw.period = jw.Interval - time.Now().Sub(started) + jw.GetJitter()

		// If the period is less than the minimum wait, then wait that long.
		if jw.period < MinimumWait {
			jw.period = MinimumWait
		}

	}
}

// Shutdown stops the forever running of the interval routine.
func (jw *JitteredInterval) Shutdown() {
	jw.Stopped = true

	jw.ShutdownChannel <- "Down"
	<-jw.ShutdownChannel

	close(jw.ShutdownChannel)
}

// GetJitter returns a positive or negative number of seconds with a mean of
// zero and a standard deviation of 2, effectivelly jittering the interval
// between around -6 and 6 seconds, though mostly not doing very much jitter.
func (jw *JitteredInterval) GetJitter() time.Duration {
	sample := rand.NormFloat64() * JitterSigma
	return time.Duration(sample*1000) * time.Millisecond
}
