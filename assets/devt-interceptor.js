// DevTool Request Interceptor
// Automatically adds devt_parent_uid header to all fetch and XMLHttpRequest calls

(function() {
    'use strict';

    // Get the devt_parent_uid from the page meta tag
    function getDevtParentUid() {
        const meta = document.querySelector('meta[name="devt-parent-uid"]');
        return meta ? meta.content : null;
    }

    // Store original fetch function
    const originalFetch = window.fetch;

    // Override global fetch function
    window.fetch = function(url, options = {}) {
        const devtParentUid = getDevtParentUid();

        if (devtParentUid) {
            // Ensure headers object exists
            if (!options.headers) {
                options.headers = {};
            }

            // Add devt_parent_uid header
            options.headers['devt_parent_uid'] = devtParentUid;
        }

        // Call original fetch with modified options
        return originalFetch.call(this, url, options);
    };

    // Store original XMLHttpRequest methods
    const originalXHROpen = XMLHttpRequest.prototype.open;
    const originalXHRSend = XMLHttpRequest.prototype.send;

    // Override XMLHttpRequest.open to store the URL
    XMLHttpRequest.prototype.open = function(method, url, async, user, password) {
        this._devtUrl = url;
        this._devtMethod = method;
        return originalXHROpen.apply(this, arguments);
    };

    // Override XMLHttpRequest.send to add the header
    XMLHttpRequest.prototype.send = function(data) {
        const devtParentUid = getDevtParentUid();

        if (devtParentUid && this.readyState === XMLHttpRequest.OPENED) {
            this.setRequestHeader('devt_parent_uid', devtParentUid);
        }

        return originalXHRSend.apply(this, arguments);
    };

    // Log initialization in dev mode
    if (window.console && console.log) {
        console.log('[DevTool] Request interceptor initialized');
    }

})();