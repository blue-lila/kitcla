function search() {
    return {
        search: "",
        results: [],
        cursor: 0,
        lastSearch: "",
        kind: 0,
        visible: false,
        searchKinds: [
            {"key":"page","text":"Page"},
            {"key":"scoped","text":"Local"},
            {"key":"global","text":"Global"},
            {"key":"app","text":"Run"},
        ],
        init() {
            window.addEventListener("_search_toggled", (event) => {
                this.onToggle()
            });
        },
        onToggle() {
            this.visible = !this.visible;
            if (this.visible) {
                this.fetchSearchResults();
            }
        },
        doSearch() {
            if (this.lastSearch === this.search) {
                return
            }
            this.lastSearch = this.search;
            this.fetchSearchResults()
        },
        fetchSearchResults() {
            let scope = document.getElementById('_search-search-component').dataset.scope;
            scope = JSON.parse(scope)
            fetch("/ajx/search", {
                method: "POST",
                body: JSON.stringify({
                    search: this.search,
                    kind: this.searchKinds[this.kind].key,
                    scope: scope,
                }),
                headers: {
                    "Content-type": "application/json; charset=UTF-8"
                }
            })  .then((response) => response.json())
                .then((json) => {
                    if (!json.items) {
                        this.results = [];
                    } else {
                        this.results = json.items;
                    }
                    this.cursor = 0;
                    this.selectCurrent();
                });
        },
        close() {
            const searchElement = document.getElementById('_search-search');
            searchElement.classList.add('hidden');
            this.results = [];
            this.search = "";
            this.cursor = 0;
            this.kind = 0;
        },
        current() {
            return this.results[this.cursor];
        },
        currentKind() {
            return this.searchKinds[this.kind];
        },
        noResults() {
            return !this.results || this.results.length === 0
        },
        selectCurrent() {
            if (this.noResults()) {
                return
            }
            for (let i in this.results) {
                this.results[i].Selected = false;
            }
            this.results[this.cursor].Selected = true;
        },
        cursorDown() {
            this.cursor++;
            if (this.cursor >= this.results.length) {
                this.cursor = 0;
            }
            this.selectCurrent();
        },
        cursorUp() {
            this.cursor--;
            if (this.cursor < 0) {
                this.cursor = this.results.length - 1;
            }
            this.selectCurrent();
        },
        nextSearchKind() {
            let kind = this.kind
            kind++;
            if (kind >= this.searchKinds.length) {
                kind = 0;
            }
            this.kind = kind;
        },
        chooseCurrent() {
            if (this.noResults()) {
                return
            }
            window.location.assign(this.current().Url);
        }
    }
}