// mkdir.js

/**
 * makdir function is dependent on the 'const cwd',
 * which is in the 'up.js' file so use this after importing 'up.js'
 */
async function mkdir() {
    let input = prompt("Enter the name for the new directory");
    if (!input) {
        return;
    }

    input = input.trim();
    const url = `/mkdir?cwd=${encodeURIComponent(cwd)}&name=${encodeURIComponent(input)}`;
    console.log("req", url);
    const res = await fetch(url, {
        method: 'POST',
    });

    if (res.ok) {
        window.location.href = `/upload?cwd=${encodeURIComponent(cwd[cwd.length - 1] === "/" ? cwd +input: cwd + "/" + input)}`;
        return;
    }
    alert("Could not crate directoy. Please input a valid fileName");
}