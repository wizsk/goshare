// main.js
function showNotShowZipDown() {
    if (filesToZipCount > 0) {
        dataZipDown.style.display = "block";
        console.log("zip down showing:", filesToZipCount);
    } else {
        dataZipDown.style.display = "none";
        console.log("zip down not showing:", filesToZipCount);
    }
}


const dataZipSelect = "data-zip-select";
const items = document.querySelectorAll(`[${dataZipSelect}]`);
const dataZipDown = document.getElementById("zip-download");
let filesToZipCount = 0;

items.forEach((itm) => {
    // if (itm.checked) filesToZipCount++;
    itm.checked = false;
});

showNotShowZipDown();


document.getElementById("clear-selection").addEventListener("click", () => {
    items.forEach((itm) => {
        if (itm.checked) {
            filesToZipCount--;
            itm.checked = false;
        }
    });

    showNotShowZipDown();
});


items.forEach((itm) => {
    itm.addEventListener("click", (ev) => {
        if (itm.checked) {
            filesToZipCount++;
        } else {
            filesToZipCount--;
        }

        showNotShowZipDown();
    });
});
