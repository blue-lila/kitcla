function showTextIntoModal(el) {
    alert(el.dataset.text);
}

function performButton() {
    return {
        performAsync(url) {
            const href = url
            fetch(href, {
                method: "POST",
                body: JSON.stringify({}),
                headers: {
                    "Content-type": "application/json; charset=UTF-8"
                }
            })  .then((response) => response.json())
                .then((json) => {
                    window.location.reload();
                });
        }
    }
}

function G_getRemoteForm(url, mod = {}) {
    fetch(url, {
        method: "GET",
    })
        .then((response) => response.text())
        .then((html) => {
            const parser = new DOMParser();
            const doc = parser.parseFromString(html, "text/html");
            const formElement = doc.querySelector('[data-w-remote-form]');
            if (!formElement) {
                console.error(`Element with selected "data-w-remote-form" has not been found in the remote DOM of ` + url);
                return;
            }
            const key = formElement.getAttribute('data-w-remote-form');
            if (!key) {
                console.error(`Element with selected "data-w-remote-form" has no value`);
                return;
            }
            const targetElement = document.querySelector(`[data-w-remote-form-area="${key}"]`);
            if (!targetElement) {
                console.error(`Element with selected "data-w-remote-form-area" has not been found in the page DOM.`);
                return;
            }
            targetElement.replaceChildren(formElement);
        })
        .catch((error) => {
            console.error("Error fetching or updating DOM:", error);
        });
}

function G_postRemoteForm(event, mod = {}) {
    event.preventDefault();

    const form = event.target;
    const action = form.action;
    const method = 'POST';
    const formData = new FormData(form);

    fetch(action, {
        method: method.toUpperCase(),
        body: formData,
        headers: {
            'X-Requested-With': 'XMLHttpRequest'
        }
    })
        .then(response => response.text())
        .then(responseText => {
            window.location.reload();
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function G_reloadHtmlNodeById(id) {
    const url = window.location.href;
    fetch(url, {
        method: "GET",
        headers: {
            "id_refresh_target": id,
        }
    })
        .then((response) => response.text())
        .then((html) => {
            const parser = new DOMParser();
            const doc = parser.parseFromString(html, "text/html");
            const newElement = doc.getElementById(id);

            if (newElement) {
                const targetElement = document.getElementById(id);
                if (targetElement) {
                    targetElement.replaceWith(newElement);
                } else {
                    console.error(`Element with id "${id}" not found in the DOM.`);
                }
            } else {
                console.error(`Element with id "${id}" not found in the response HTML.`);
            }
        })
        .catch((error) => {
            console.error("Error fetching or updating DOM:", error);
        });
}

function copyToClipboard(text) {
    navigator.clipboard.writeText(text)
        .then(() => {
            // Noop for now
        })
        .catch(err => {
            console.error('Failed to copy text: ', err);
        });
}

function debugOpenFilePath(filepath) {
    const url = "/open-file";
    fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ filepath: filepath }),
    })
        .then(response => {
            if (!response.ok) {
                throw new Error("Failed to open file");
            }
            return response.json();
        })
        .then(data => {
            console.log("File opened successfully:", data);
        })
        .catch(error => {
            console.error("Error:", error);
        });
}

function showPanel(key) {
    const panels = document.querySelectorAll('div[data-panel-key]');

    panels.forEach(panel => {
        if (panel.getAttribute('data-panel-key') === key) {
            panel.classList.remove('hidden');
        } else {
            panel.classList.add('hidden');
        }
    });
}

document.addEventListener('keydown', function(event) {
    if (event.ctrlKey && event.key === 'k') {
        event.preventDefault();
        const searchElement = document.getElementById('_search-search');
        const searchInput = document.getElementById('_search-search-input');

        if (searchElement) {
            const event = new CustomEvent("_search_toggled");
            window.dispatchEvent(event);
            searchElement.classList.toggle('hidden');
            if (!searchElement.classList.contains('hidden') && searchInput) {
                searchInput.focus();
            }
        }
    }
});
