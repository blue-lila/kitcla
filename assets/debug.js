function debug(port) {
    return {
        retryCount: 0,
        beenConnected: false,
        url: 'http://localhost:' + port + '/sse',
        lostConnection: false,
        showIndicator: false,
        eventSource: null,
        connect() {
            this.eventSource = new EventSource(this.url);

            this.eventSource.onopen = (event) => {
                debugInstance = this;
                if (this.lostConnection && this.beenConnected) {
                    window.location.reload();
                }
                this.beenConnected = true;
                this.retryCount = 0;
                this.showIndicator = false;
            };

            this.eventSource.onmessage = (event) => {
                //console.log("Debug message:", event.data);
            };

            this.eventSource.onerror = (event) => {
                console.log("SSE Debug connection closed");
                this.lostConnection = true;

                if (this.retryCount > 0) {
                    this.showIndicator = true;
                }

                if (this.retryCount > 10) {
                    console.log("Max retry attempts reached");
                    return;
                }

                this.retryCount++;
                setTimeout(() => this.connect(), 500 * this.retryCount);
            };
        },
        close() {
            if (this.eventSource) {
                this.eventSource.close();
                this.eventSource = null;
            }
        }
    }
}

function xGocDebugElement() {
    console.log("Printing Goc debug")
    let nodeList = document.querySelectorAll(":hover");
    let ids = Array.from(nodeList).filter(a => a.dataset.gxi).map(a => a.dataset.gxi);
    printXGocDebugElement(ids[ids.length - 1]);
}


function printXGocDebugElement(id) {
    fetch('http://localhost:9645/open-gix?id=' + id)
        .catch(error => console.error('Error:', error));
    console.log(id);
}

// Global debug instance for cleanup
let debugInstance = null;

window.onkeydown = evt => {
    switch (evt.keyCode) {
        // F9 key
        case 120:
            xGocDebugElement();
            break;
        default:
            return true;
    }
    //Returning false overrides default browser event
    return false;
};

// Close SSE connection before page unload
window.addEventListener('beforeunload', () => {
    if (debugInstance) {
        debugInstance.close();
    }
});