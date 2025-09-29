function advancedInput(initialValue) {
    return {
        initialValue: initialValue,
        filter: '',
        show: false,
        selected: null,
        value: "",
        focusedOptionIndex: null,
        options: null,
        close() {
            this.show = false;
            this.filter = this.selectedName();
            this.focusedOptionIndex = this.selected ? this.focusedOptionIndex : null;
        },
        open() {
            this.show = true;
            this.filter = '';
        },
        toggle() {
            if (this.show) {
                this.close();
            } else {
                this.open()
            }
        },
        isOpen() {
            return this.show === true
        },
        selectedName() {
            return this.selected ? this.selected.text : this.filter;
        },
        classOption(id, index) {
            const isSelected = this.selected ? (id == this.selected.id) : false;
            const isFocused = (index == this.focusedOptionIndex);
            return {
                'cursor-pointer w-full border-gray-100 border-b hover:bg-blue-50': true,
                'bg-blue-100': isSelected,
                'bg-blue-50': isFocused
            };
        },
        fetchOptions(key, filters) {
            if (key === "__test__") {
                this.options = {items: [{id: 1, text: "placeholder 1"}, {id: 2, text: "placeholder 2"}]}
                return
            }

            fetch('/ajx/autocomplete', {
                method: 'POST',
                body: JSON.stringify({key: key, search: "", filters: filters}),
            })
                .then(response => response.json())
                .then(data => {
                    this.options = data;
                    this.selectOptionFromInitial();
                });
        },
        selectOptionFromInitial() {
            if (!this.options.items) {
                return;
            }

            let selected = this.options.items.find(option => option.id === this.initialValue);
            if (!selected) {
                return;
            }
            console.log(selected);
            this.selected = selected;
            this.filter = this.selectedName();
            this.value = selected.id;
        },
        filteredOptions() {
            if (!this.options) {
                return;
            }
            return this.options && this.options.items
                ? this.options.items.filter(option => {
                    return (option.text.toLowerCase().indexOf(this.filter) > -1)
                })
                : {}
        },
        onOptionClick(index) {
            this.focusedOptionIndex = index;
            this.selectOption();
        },
        selectOption() {
            if (!this.isOpen()) {
                return;
            }
            this.focusedOptionIndex = this.focusedOptionIndex ?? 0;
            const selected = this.filteredOptions()[this.focusedOptionIndex]
            if (this.selected && this.selected.id == selected.id) {
                this.filter = '';
                this.selected = null;
                this.value = "";
            } else {
                this.selected = selected;
                this.filter = this.selectedName();
                this.value = selected.id;
            }
            this.close();
        },
        focusPrevOption() {
            if (!this.isOpen()) {
                return;
            }
            const optionsNum = Object.keys(this.filteredOptions()).length - 1;
            if (this.focusedOptionIndex > 0 && this.focusedOptionIndex <= optionsNum) {
                this.focusedOptionIndex--;
            } else if (this.focusedOptionIndex == 0) {
                this.focusedOptionIndex = optionsNum;
            }
        },
        focusNextOption() {
            const optionsNum = Object.keys(this.filteredOptions()).length - 1;
            if (!this.isOpen()) {
                this.open();
            }
            if (this.focusedOptionIndex == null || this.focusedOptionIndex == optionsNum) {
                this.focusedOptionIndex = 0;
            } else if (this.focusedOptionIndex >= 0 && this.focusedOptionIndex < optionsNum) {
                this.focusedOptionIndex++;
            }
        }
    }
}

function form() {
    return {
        redirect: "",
    }
}