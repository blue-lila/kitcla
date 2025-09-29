function table(tableId) {
    return {
        tableId: tableId,
        showFilter: false,
        showSelectColumns: false,
        filterKey: "",
        operationKey: "",
        value: "",
        toggleFiltersChoice() {
            if (this.showFilter === false) {
                this.value = "";
                const select = document.getElementById("_new_filter");
                this.filterKey = select.options[0].value;
                const select2 = document.getElementById("_new_filter_" + this.filterKey);
                this.operationKey = select2.options[0].value;
            }
            this.showFilter = !this.showFilter
        },
        toggleSelectColumns() {
            this.showSelectColumns = !this.showSelectColumns;
        },
        selectColumn(key) {
            let url = new URL(window.location);
            let params = new URLSearchParams(url.search);
            let name = 'columns[visibility][' + key + "]"
            params.set(name, true);
            window.location.href = url.pathname + '/?' + params;
        },
        chooseFilter(event) {
            this.value = "";
            this.filterKey = event.target.value;
            const id = "_new_filter_" + this.filterKey
            const select = document.getElementById(id);
            this.operationKey = select.optionss[0].value;
        },
        chooseOperation(event) {
            this.value = "";
            this.operationKey = event.target.value;
        },
        createFilter() {
            let url = new URL(window.location);
            let params = new URLSearchParams(url.search);
            let name = 'filters[fields][' + this.filterKey + "][" + this.operationKey + "]"
            params.set(name, this.value);
            window.location.href = url.pathname + '/?' + params;
        },
        removeFilter(filterKey, operationKey) {
            let url = new URL(window.location);
            let params = new URLSearchParams(url.search);
            let name = 'filters[fields][' + filterKey + "][" + operationKey + "]"
            params.delete(name);
            window.location.href = url.pathname + '/?' + params;
        },
        checkAllIds() {
            let controller = document.querySelector('[data-table-type="checkbox_ids_controller"][data-table-id="' + this.tableId + '"]')
            let items = document.querySelectorAll('[data-table-type="checkbox_id"][data-table-id="' + this.tableId + '"]')
            items.forEach((item) => item.checked = controller.checked)
        },
        collectCheckedIds() {
            const selector = 'input[type="checkbox"]:checked[data-table-type="checkbox_id"][data-table-id="' + this.tableId + '"]'
            let items = Array.from(document.querySelectorAll(selector))
            return items.map((x) => x.dataset.id);
        },
        submitCheckedIds(el) {
            const href = el.dataset.href
            const ids = this.collectCheckedIds()
            fetch(href, {
                method: "POST",
                body: JSON.stringify({
                    ids: ids,
                }),
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