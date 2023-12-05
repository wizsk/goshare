// main.js

window.onload = () => {
    // runs every time to clear the selections
    items.forEach((itm) => {
        filesToZipCount = 0;
        itm.checked = false;
    });
    showNotShowZipDown();
}

const dataZipSelect = "data-zip-select";
const items = document.querySelectorAll(`[${dataZipSelect}]`);
const dataZipDown = document.getElementById("zip-download");
let filesToZipCount = 0;

function showNotShowZipDown() {
    if (filesToZipCount > 0) {
        dataZipDown.style.display = "block";
        console.log("zip down showing:", filesToZipCount);
    } else {
        dataZipDown.style.display = "none";
        console.log("zip down not showing:", filesToZipCount);
    }
}



function selectAll() {
    items.forEach((itm) => {
        itm.checked = true;
    });
    filesToZipCount = items.length;
    showNotShowZipDown();
}


function clearSelections() {
    items.forEach((itm) => {
        itm.checked = false;
    });

    filesToZipCount = 0;
    showNotShowZipDown();
}


function markUnmark(itm) {
    if (itm.checked) {
        filesToZipCount++;
    } else {
        filesToZipCount--;
    }
    showNotShowZipDown();
}
