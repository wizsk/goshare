// zip.js

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
        a.href = data.url;
        a.style.display = 'none';
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);
        zipDownProgress.innerText = e.data;
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

    // fetch(fo);
}