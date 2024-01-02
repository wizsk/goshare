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

/**
 * zipOptionsSH is the html btn for 
 * @type {HTMLElement}
 */
const zipOptionsSH = document.getElementById("zip-options");

/**
 * @type {HTMLElement}
 */
const optionsBackdrop = document.getElementById("opt-backdrop");

function showHideZipOptions() {
    if (window.screen.width > 768) return;
    

    if (!document.getElementById("zip-options").classList.toggle("hidden")) {
        optionsBackdrop.classList.remove("hidden");
    } else {
        optionsBackdrop.classList.add("hidden");
    }
}

optionsBackdrop.addEventListener("click", () => {
    zipOptionsSH.classList.add("hidden")
    optionsBackdrop.classList.add("hidden")
});


// document.addEventListener("click", () => {
//     if (!zipOptionsSH.classList.contains("hidden")) {
//         zipOptionsSH.classList.add("hidden");
//         return;
//     }
// })


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
    showHideZipOptions();
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
