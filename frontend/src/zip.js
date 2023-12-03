
// zip.js

const zipDownBtn = document.getElementById("zip-download");
let isZippin = false;

zipDownBtn.addEventListener("click", () => {
    if (isZippin) return;
    zipDownBtn.disabled = true;
    isZippin = true;

    const url = [];
    items.forEach((itm) => {
        if (itm.checked) {
            url.push(itm.getAttribute(dataZipSelect));
        }
    });

    const strr = url.map(itm => `files=${encodeURIComponent(itm)}`).join("&");
    const path = `/zip?${strr}`; //&files=${encodeURIComponent("/../../../")}`;

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
        sse.close();
    });

    sse.addEventListener("onProgress", async (e) => {
        console.log("sse onPgoress", e.data);
    });

    sse.onclose = () => {
        zipDownBtn.disabled = false;
        isZippin = false;
        console.log("SSE connection closed.");
    };

    // fetch(fo);
});