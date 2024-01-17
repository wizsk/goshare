// mkdir.js

// const mkdirform = document.getElementById("mkdir-form");
// const input = document.getElementById("mkdir-input");


// mkdirform.addEventListener("submit", async (e) => {
//     e.preventDefault();
// })


async function mkdir() {
    let input = prompt("enter the name for the directory");
    if (!input) {
        return;
    }

    const url = `/mkdir?cwd=${encodeURIComponent(window.location.pathname)}&name=${encodeURIComponent(input)}`;
    console.log("req", url);
    const res = await fetch(url, {
        method: 'POST',
    });
    if (res.ok && res.redirected) {
        window.location.href = res.url;
        return;
    }
    alert("Please input a valid fileName");
    mkdir();
}