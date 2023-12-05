// mkdir.js

const mkdirform = document.getElementById("mkdir-form");
const input = document.getElementById("mkdir-input");


mkdirform.addEventListener("submit", async (e) => {
    e.preventDefault();
    const url = `/mkdir?cwd=${encodeURIComponent(window.location.pathname)}&name=${encodeURIComponent(input.value)}`;
    console.log("req", url);
    const res = await fetch(url, {
        method: 'POST',
    });
    if (res.ok && res.redirected) {
        window.location.href = res.url;
    }
})