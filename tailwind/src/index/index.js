// {{define "index-js"}}
const body = document.body;
const light =
    '<svg class="w-full h-full fill-inherit" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><title>brightness-6</title><path d="M12,18V6A6,6 0 0,1 18,12A6,6 0 0,1 12,18M20,15.31L23.31,12L20,8.69V4H15.31L12,0.69L8.69,4H4V8.69L0.69,12L4,15.31V20H8.69L12,23.31L15.31,20H20V15.31Z" /></svg>';
const dark =
    '<svg class="w-full h-full fill-inherit" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><title>brightness-4</title><path d="M12,18C11.11,18 10.26,17.8 9.5,17.45C11.56,16.5 13,14.42 13,12C13,9.58 11.56,7.5 9.5,6.55C10.26,6.2 11.11,6 12,6A6,6 0 0,1 18,12A6,6 0 0,1 12,18M20,8.69V4H15.31L12,0.69L8.69,4H4V8.69L0.69,12L4,15.31V20H8.69L12,23.31L15.31,20H20V15.31L23.31,12L20,8.69Z" /></svg>';

const theme_btn = document.getElementById("theme-btn");
const menu = document.getElementById("menu");
const menu_items = document.getElementById("menu-items");
const menu_cancel = document.getElementById("menu-cancel");
const zip_file_divs = document.querySelectorAll(".zip-file");
const search_bar = document.getElementById("search");
let Ziping = false;

// On page load or when changing themes, best to add inline in to avoid FOUC
if (
    localStorage.theme === "dark" ||
    (!("theme" in localStorage) &&
        window.matchMedia("(prefers-color-scheme: dark)").matches)
) {
    body.classList.add("dark");
    theme_btn.innerHTML = dark;
} else {
    body.classList.remove("dark");
    theme_btn.innerHTML = light;
}


menu.addEventListener("click", () => {
    menu.style.display = "none";
    menu_items.classList.remove("hidden");
});

menu_cancel.addEventListener("click", () => {
    menu.style.display = "block";
    menu_items.classList.add("hidden");
});

theme_btn.addEventListener("click", () => {
    if (body.classList.contains("dark")) {
        body.classList.remove("dark");
        theme_btn.innerHTML = light;
        localStorage.theme = "light";
    } else {
        body.classList.add("dark");
        theme_btn.innerHTML = dark;
        localStorage.theme = "dark";
    }
});

const file_name_rows = document.querySelectorAll(".file-name-rows");
search_bar.addEventListener("input", (e) => {
    let value = e.target.value.toLowerCase();
    file_name_rows.forEach((file) => {
        const visible = file
            .querySelector(".file-names")
            .innerText.toLowerCase()
            .includes(value);

        file.classList.toggle("hidden", !visible);
    });
});

addEventListener("keydown", (event) => {
    if (event.key === "/") {
        event.preventDefault();
        search_bar.focus();
    }
});



document.getElementById("sort").addEventListener("change", (e) => {
    let params = `${window.location.search}`;
    let cleaned_params = "";

    if (params.includes("sort")) {
        const param = params.split("&");
        for (let i = 0; i < param.length; i++) {
            itm = param[i];
            if (itm.includes("sort")) {
                continue;
            }
            if (itm !== "") {
                cleaned_params += (cleaned_params.length !== 0 ? "&" : "?") + itm;
            }
        }
    }
    cleaned_params +=
        cleaned_params.length > 0
            ? "&sort=" + e.target.value
            : "?sort=" + e.target.value;

    const fullURL = `${window.location.protocol}//${window.location.host}${window.location.pathname}${cleaned_params}${window.location.hash}`;
    window.location.assign(fullURL);
});

let progress_div = document.getElementById("progress-div");
function eventM(link) {
    console.log("sending sse request to to for sse:", link);
    progress_div.classList.remove("hidden");
    const eventSource = new EventSource(link);

    // Event listener to handle errors and closed connections
    eventSource.onerror = function (event) {
        console.error("Error occurred:", event);
        eventSource.close();
    };

    // castom err event
    eventSource.addEventListener("error", async (e) => {
        progress_div.innerText = "cond't zip"
        eventSource.close();
        Ziping = false;
    });

    // done event
    eventSource.addEventListener("done", async (e) => {
        const jsn = await JSON.parse(e.data);
        console.log("zipping done:", jsn.name);
        // Create a new <a> element
        const link = document.createElement('a');
        link.classList.add("block", "p-3", "hover:underline")
        link.href = `${jsn.url}?zip=down`;
        // link.download = jsn.status;
        link.innerText = `Download ${jsn.name}`;

        // Append the <a> element to the specified container
        // zip_progress.innerHTML = `<a href="${jsn.status}?zip=down">Download zip</a>`;
        progress_div = zip_progress.removeChild(progress_div);
        zip_progress.appendChild(link);
        eventSource.close();
        Ziping = false;
    });

    eventSource.addEventListener("onProgress", async (e) => {
        if (!document.getElementById("progress-div")) {
            zip_progress.appendChild(progress_div);
        }
        const jsn = await JSON.parse(e.data);
        progress_div.innerText = jsn.status;
    });

    eventSource.onclose = function () {
        console.log("SSE connection closed.");
    };
}

const zip_progress = document.getElementById("zip-progress");
function getZipFile(elemnt) {
    if (Ziping) return;
    Ziping = true;
    const link = elemnt.getAttribute("data-link");
    eventM(link);
}

zip_file_divs.forEach((elemnt) => {
    elemnt.addEventListener("click", (e) => {
        console.log("getting stuff:", elemnt)
        getZipFile(elemnt);
        elemnt.classList.remove("zip-file")
        const cp = elemnt.cloneNode(true);
        elemnt.parentNode.replaceChild(cp, elemnt);
    })
})

// {{end}}