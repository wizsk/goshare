// main.js

window.onload = () => {
    // runs every time to clear the selections
    items.forEach((itm) => {
        itm.checked = false;
    });
    filesToZipCount = 0;
    showNotShowZipDown();
}

const dataZipSelect = "data-zip-select";
const items = document.querySelectorAll(`[${dataZipSelect}]`);
const dataZipDown = document.getElementById("zip-download");
// const zipSelectAllBtn = document.getElementById("selectAll-button");
const zipClearSelectionBtn = document.getElementById("clearSelections-button");
let filesToZipCount = 0;

/** shows or hides downzip button, zipClearSelections buttton */
function showNotShowZipDown() {
    if (filesToZipCount > 0) {
        dataZipDown.style.display = "block";
        zipClearSelectionBtn.style.display = "block";
        console.log("zip down showing:", filesToZipCount);
    } else {
        dataZipDown.style.display = "none";
        zipClearSelectionBtn.style.display = "none";
        console.log("zip down not showing:", filesToZipCount);
    }
}

function showHideZipOptions() {
    document.getElementById("zip-options").classList.toggle("hidden");
}


/** selectAll function selets all the file in current direcotr for zipping */
function selectAll() {
    items.forEach((itm) => {
        itm.checked = true;
    });
    filesToZipCount = items.length;
    showNotShowZipDown();
}


/** clearSelections function deselets all the file in current direcotr for zipping */
function clearSelections() {
    items.forEach((itm) => {
        itm.checked = false;
    });

    filesToZipCount = 0;
    showNotShowZipDown();
}


/**
 * @param {HTMLInputElement} itm - htmlInputElelemt type checkBox
 */
function markUnmark(itm) {
    if (itm.checked) {
        filesToZipCount++;
    } else {
        filesToZipCount--;
    }
    showNotShowZipDown();
}
