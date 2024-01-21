// main.js

const dataZipSelect = "data-zip-select";
const items = document.querySelectorAll(`[${dataZipSelect}]`);

/**
 * zipOptionsSH is the html btn for 
 * @type {HTMLElement}
 */
const zipOptionsSH = document.getElementById("zip-options");

/**
 * @type {HTMLElement}
 */
const backdrop = document.getElementById("backdrop");

function showHideZipOptions() {
    if (window.screen.width > 768) return;// md: 768px

    if (!document.getElementById("zip-options").classList.toggle("hidden")) {
        document.body.style.overflow = "hidden";
        backdrop.classList.remove("hidden");
    } else {
        backdrop.classList.add("hidden");
        document.body.style.overflow = "";
    }
}

backdrop.addEventListener("click", () => {
    zipOptionsSH.classList.add("hidden")
    backdrop.classList.add("hidden")
    document.body.style.overflow = "";
});


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
    zipDownBtn.disabled = true;
    isZippin = true;


    const strr = url.map(itm => `files=${encodeURIComponent(itm)}`).join("&");
    const path = `/zip?${strr}&cwd=${encodeURIComponent(window.location.pathname)}`; //&files=${encodeURIComponent("/../../../")}`;

    console.log("SSE reqest to:", path);
    const sse = new EventSource(path);

    sse.onopen = () => {
        zipDownBtn.disabled = true;
        zipDownProgress.classList.remove("hidden");
        isZippin = true;
        console.log("sse opended");
    }

    sse.onerror = (err) => {
        zipDownBtn.disabled = false;
        isZippin = false;
        console.error(err);
        sse.close();
    }

    // done event
    sse.addEventListener("done", async (e) => {
        zipDownBtn.disabled = false;
        isZippin = false;
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
        sse.close();
    });

    sse.addEventListener("onProgress", async (e) => {
        zipDownBtn.classList.remove("hidden");
        zipDownProgress.innerText = e.data;
        const data = JSON.parse(e.data);
        console.log("sse onPgoress", e.data);
        zipDownBtn.innerText = `Zipping: ${data.status}%`;
    });

    sse.onclose = () => {
        zipDownBtn.disabled = false;
        isZippin = false;
        console.log("SSE connection closed.");
    };
}