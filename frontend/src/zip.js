// main.js

const dataZipSelect = "data-zip-select";
const items = document.querySelectorAll(`[${dataZipSelect}]`);


/**
 * zipOptions is the html btn for
 * @type {HTMLElement}
 */
const zipOptions = document.getElementById("zip-options");

/**
 * @type {HTMLElement}
 */
const backdrop = document.getElementById("backdrop");

backdrop.addEventListener("click", () => {
    zipOptions.classList.add("hidden")
    backdrop.classList.add("hidden")
    document.body.style.overflow = "";
});

function showHideZipOptions() {
    // dont run in a bigger screen
    if (window.screen.width > 768) return;

    if (!zipOptions.classList.toggle("hidden")) {
        document.body.style.overflow = "hidden";
        backdrop.classList.remove("hidden");
    } else {
        backdrop.classList.add("hidden");
        document.body.style.overflow = "";
    }
}


/** selectAll function selets all the file in current direcotr for zipping */
function selectAll() {
    items.forEach((itm) => {
        itm.checked = true;
    });
}


/** clearSelections function deselets all the file in current direcotr for zipping */
function clearSelections() {
    items.forEach((itm) => {
        itm.checked = false;
    });
    showHideZipOptions();
}

window.onload = () => items.forEach((itm) => { itm.checked = false; });;

//////////////////
// zip ssr
//////////////////

const zipDownProgress = document.getElementById("zip-down-progress");
const zipDownBtn = document.getElementById("zip-download");
let isZippin = false;

function downloadAsZip() {
    const url = [];
    items.forEach((itm) => {
        if (itm.checked) {
            url.push(itm.getAttribute(dataZipSelect));
        }
    });

    if (url.length === 0) {
        alert("Please select some files");
        return;
    }
    showHideZipOptions();

    if (isZippin) {
        alert("Already zipping");
        return
    }


    const strr = url.map(itm => `files=${encodeURIComponent(itm)}`).join("&");
    const path = `/zip?${strr}&cwd=${encodeURIComponent(window.location.pathname)}`; //&files=${encodeURIComponent("/../../../")}`;

    console.log("SSE reqest to:", path);
    const sse = new EventSource(path);

    sse.onopen = () => {
        zipDownBtn.disabled = true;
        isZippin = true;
        zipDownProgress.classList.remove("hidden");
        zipDownProgress.innerText = "Started zipping";
        console.log("sse opended");
    }

    sse.onerror = (err) => {
        console.error(err);
        zipDownProgress.innerText = "Something went wrong";
        sse.close();
        zipDownBtn.disabled = false;
        isZippin = false;
    }

    sse.addEventListener("restricted", async (e) => {
        zipDownProgress.innerText = 'Zipping not allowed';
        sse.close();
        zipDownBtn.disabled = false;
        isZippin = false;
    });

    // done event
    sse.addEventListener("done", async (e) => {
        console.log("sse done:", e.data);
        const data = JSON.parse(e.data);
        const a = document.createElement("a");

        a.innerText = `Download: ${data.name}`;
        a.download = data.name;
        a.href = data.url;

        a.classList.add("hover:underline")

        a.click();

        zipDownProgress.innerText = "";
        zipDownProgress.appendChild(a);

        zipDownBtn.disabled = false;
        isZippin = false;
        sse.close();
    });

    sse.addEventListener("onProgress", async (e) => {
        const data = JSON.parse(e.data);
        console.log("sse onPgoress", data.status);
        zipDownProgress.innerText = `Zipping: ${data.status}%`;
    });

    sse.onclose = () => {
        zipDownBtn.disabled = false;
        isZippin = false;
        console.log("SSE connection closed.");
    };
}
