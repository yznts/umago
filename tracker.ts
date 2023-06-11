

// Please note, that this is just a simple example of a tracker.
// It is not intended to be used in production,
// Google doesn't like open trackers.
// Feel free to use it as a base for your own one.
class Tracker {

    // Server may register a tracking handler anywhere,
    // so we need to know that endpoint.
    endpoint: string

    // We are initializing tracker with a tracking endpoint.
    constructor(endpoint: string) {
        // We are saving endpoint to use it later.
        this.endpoint = endpoint
        // Bind custom events tracking.
        // Use `data-event="event-name"` attribute to track custom events.
        document.querySelectorAll(`[data-event]`).forEach(el => {
            el.addEventListener('click', () => {
                // Extract event name
                let event = el.getAttribute('data-event')
                // Pass, if no event name
                if (!event) {
                    return
                }
                // Track event
                this.send(event)
            })
        })
    }

    // Track a new event.
    // The only optional parameter, event name,
    // allows you to track custom events.
    // If event name is not provided, it's a page view.
    send(event?: string) {
        fetch(this.endpoint, {
            method: 'POST',
            body: JSON.stringify({
                n: event,
                t: document.title,
                r: document.referrer,
            })
        })
    }
}
