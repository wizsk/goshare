package main

import "html/template"

var form = []byte(`<!DOCTYPE html> <html> <head> <meta charset="utf-8" /> <meta name="viewport" content="width=device-width, initial-scale=1.0" /> <title>auth</title> <style> *, *:before, *:after { -moz-box-sizing: border-box; -webkit-box-sizing: border-box; box-sizing: border-box; } body { font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif; color: #384047; } form { max-width: 90%; margin: 10px auto; padding: 10px 20px; background: #f4f7f8; border-radius: 8px; margin-top: 20vh; } h1 { margin: 0 0 30px 0; text-align: center; } input[type="password"] { background: rgba(255, 255, 255, 0.1); border: none; font-size: 16px; height: auto; margin: 0; outline: 0; padding: 0.7rem; width: 100%; background-color: #e8eeef; color: #8a97a0; box-shadow: 0 1px 0 rgba(0, 0, 0, 0.03) inset; margin-bottom: 30px; } button { padding: 0.7rem 1rem; color: #fff; background-color: #4bc970; font-size: 18px; text-align: center; font-style: normal; border-radius: 5px; width: 100%; border: 1px solid #3ac162; border-width: 1px 1px 3px; box-shadow: 0 -1px 0 rgba(255, 255, 255, 0.1) inset; margin-bottom: 10px; } fieldset { margin-bottom: 30px; border: none; } legend { font-size: 1.4em; margin-bottom: 10px; } label { display: block; margin-bottom: 8px; } label.light { font-weight: 300; display: inline; } @media screen and (min-width: 480px) { form { max-width: 480px; } } @media (prefers-color-scheme: dark) { body { background-color: black; color: white; } form { background-color: #384047; } input[type="password"] { background-color: #4c535a; /* background: #3ac162; */ } } </style> </head> <body> <div> <div> <form action="/dQw4w9WgXcQ/auth" method="post"> <h1>Authorize</h1> <fieldset> <label for="password">Password:</label> <input autofocus type="password" id="password" name="password" /> <button type="submit">Go</button> </fieldset> </form> </div> </div> </body> </html>`)

var directoryIcon template.HTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-folder2-open" viewBox="0 0 16 16"> <path d="M1 3.5A1.5 1.5 0 0 1 2.5 2h2.764c.958 0 1.76.56 2.311 1.184C7.985 3.648 8.48 4 9 4h4.5A1.5 1.5 0 0 1 15 5.5v.64c.57.265.94.876.856 1.546l-.64 5.124A2.5 2.5 0 0 1 12.733 15H3.266a2.5 2.5 0 0 1-2.481-2.19l-.64-5.124A1.5 1.5 0 0 1 1 6.14V3.5zM2 6h12v-.5a.5.5 0 0 0-.5-.5H9c-.964 0-1.71-.629-2.174-1.154C6.374 3.334 5.82 3 5.264 3H2.5a.5.5 0 0 0-.5.5V6zm-.367 1a.5.5 0 0 0-.496.562l.64 5.124A1.5 1.5 0 0 0 3.266 14h9.468a1.5 1.5 0 0 0 1.489-1.314l.64-5.124A.5.5 0 0 0 14.367 7H1.633z"/> </svg>`

// const downloadIcon template.HTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-file-earmark-arrow-down-fill" viewBox="0 0 16 16"> <path d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0zM9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1zm-1 4v3.793l1.146-1.147a.5.5 0 0 1 .708.708l-2 2a.5.5 0 0 1-.708 0l-2-2a.5.5 0 0 1 .708-.708L7.5 11.293V7.5a.5.5 0 0 1 1 0z"/> </svg>`

const imgIcon template.HTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-file-earmark-image-fill" viewBox="0 0 16 16"> <path d="M4 0h5.293A1 1 0 0 1 10 .293L13.707 4a1 1 0 0 1 .293.707v5.586l-2.73-2.73a1 1 0 0 0-1.52.127l-1.889 2.644-1.769-1.062a1 1 0 0 0-1.222.15L2 12.292V2a2 2 0 0 1 2-2zm5.5 1.5v2a1 1 0 0 0 1 1h2l-3-3zm-1.498 4a1.5 1.5 0 1 0-3 0 1.5 1.5 0 0 0 3 0z"/> <path d="M10.564 8.27 14 11.708V14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-.293l3.578-3.577 2.56 1.536 2.426-3.395z"/> </svg>`

const videoIcon template.HTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-file-play-fill" viewBox="0 0 16 16"> <path d="M12 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zM6 5.883a.5.5 0 0 1 .757-.429l3.528 2.117a.5.5 0 0 1 0 .858l-3.528 2.117a.5.5 0 0 1-.757-.43V5.884z"/> </svg>`

const audioIcon template.HTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-file-music-fill" viewBox="0 0 16 16"> <path d="M12 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm-.5 4.11v1.8l-2.5.5v5.09c0 .495-.301.883-.662 1.123C7.974 12.866 7.499 13 7 13c-.5 0-.974-.134-1.338-.377-.36-.24-.662-.628-.662-1.123s.301-.883.662-1.123C6.026 10.134 6.501 10 7 10c.356 0 .7.068 1 .196V4.41a1 1 0 0 1 .804-.98l1.5-.3a1 1 0 0 1 1.196.98z"/> </svg>`

const pdfIcon template.HTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-file-earmark-pdf-fill" viewBox="0 0 16 16"> <path d="M5.523 12.424c.14-.082.293-.162.459-.238a7.878 7.878 0 0 1-.45.606c-.28.337-.498.516-.635.572a.266.266 0 0 1-.035.012.282.282 0 0 1-.026-.044c-.056-.11-.054-.216.04-.36.106-.165.319-.354.647-.548zm2.455-1.647c-.119.025-.237.05-.356.078a21.148 21.148 0 0 0 .5-1.05 12.045 12.045 0 0 0 .51.858c-.217.032-.436.07-.654.114zm2.525.939a3.881 3.881 0 0 1-.435-.41c.228.005.434.022.612.054.317.057.466.147.518.209a.095.095 0 0 1 .026.064.436.436 0 0 1-.06.2.307.307 0 0 1-.094.124.107.107 0 0 1-.069.015c-.09-.003-.258-.066-.498-.256zM8.278 6.97c-.04.244-.108.524-.2.829a4.86 4.86 0 0 1-.089-.346c-.076-.353-.087-.63-.046-.822.038-.177.11-.248.196-.283a.517.517 0 0 1 .145-.04c.013.03.028.092.032.198.005.122-.007.277-.038.465z"/> <path fill-rule="evenodd" d="M4 0h5.293A1 1 0 0 1 10 .293L13.707 4a1 1 0 0 1 .293.707V14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2zm5.5 1.5v2a1 1 0 0 0 1 1h2l-3-3zM4.165 13.668c.09.18.23.343.438.419.207.075.412.04.58-.03.318-.13.635-.436.926-.786.333-.401.683-.927 1.021-1.51a11.651 11.651 0 0 1 1.997-.406c.3.383.61.713.91.95.28.22.603.403.934.417a.856.856 0 0 0 .51-.138c.155-.101.27-.247.354-.416.09-.181.145-.37.138-.563a.844.844 0 0 0-.2-.518c-.226-.27-.596-.4-.96-.465a5.76 5.76 0 0 0-1.335-.05 10.954 10.954 0 0 1-.98-1.686c.25-.66.437-1.284.52-1.794.036-.218.055-.426.048-.614a1.238 1.238 0 0 0-.127-.538.7.7 0 0 0-.477-.365c-.202-.043-.41 0-.601.077-.377.15-.576.47-.651.823-.073.34-.04.736.046 1.136.088.406.238.848.43 1.295a19.697 19.697 0 0 1-1.062 2.227 7.662 7.662 0 0 0-1.482.645c-.37.22-.699.48-.897.787-.21.326-.275.714-.08 1.103z"/> </svg>`

const textIcon template.HTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-file-earmark-text-fill" viewBox="0 0 16 16"> <path d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0zM9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1zM4.5 9a.5.5 0 0 1 0-1h7a.5.5 0 0 1 0 1h-7zM4 10.5a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5zm.5 2.5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1h-4z"/> </svg>`

const unknownFileIcon template.HTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-file-earmark-minus-fill" viewBox="0 0 16 16"> <path d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0zM9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1zM6 8.5h4a.5.5 0 0 1 0 1H6a.5.5 0 0 1 0-1z"/> </svg>`

const archiveIcon template.HTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-file-earmark-zip-fill" viewBox="0 0 16 16"> <path d="M5.5 9.438V8.5h1v.938a1 1 0 0 0 .03.243l.4 1.598-.93.62-.93-.62.4-1.598a1 1 0 0 0 .03-.243z"/> <path d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0zM9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1zm-4-.5V2h-1V1H6v1h1v1H6v1h1v1H6v1h1v1H5.5V6h-1V5h1V4h-1V3h1zm0 4.5h1a1 1 0 0 1 1 1v.938l.4 1.599a1 1 0 0 1-.416 1.074l-.93.62a1 1 0 0 1-1.109 0l-.93-.62a1 1 0 0 1-.415-1.074l.4-1.599V8.5a1 1 0 0 1 1-1z"/> </svg>`

const documentIcon template.HTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-file-earmark-fill" viewBox="0 0 16 16"> <path d="M4 0h5.293A1 1 0 0 1 10 .293L13.707 4a1 1 0 0 1 .293.707V14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2zm5.5 1.5v2a1 1 0 0 0 1 1h2l-3-3z"/> </svg>`

var favicon = []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-folder-symlink-fill" viewBox="0 0 16 16"> <path d="M13.81 3H9.828a2 2 0 0 1-1.414-.586l-.828-.828A2 2 0 0 0 6.172 1H2.5a2 2 0 0 0-2 2l.04.87a1.99 1.99 0 0 0-.342 1.311l.637 7A2 2 0 0 0 2.826 14h10.348a2 2 0 0 0 1.991-1.819l.637-7A2 2 0 0 0 13.81 3zM2.19 3c-.24 0-.47.042-.683.12L1.5 2.98a1 1 0 0 1 1-.98h3.672a1 1 0 0 1 .707.293L7.586 3H2.19zm9.608 5.271-3.182 1.97c-.27.166-.616-.036-.616-.372V9.1s-2.571-.3-4 2.4c.571-4.8 3.143-4.8 4-4.8v-.769c0-.336.346-.538.616-.371l3.182 1.969c.27.166.27.576 0 .742z"/> </svg>`)

const indexPage = `
{{ define "main"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>{{.Dirtitle}}</title>

    <style>
    /*! tailwindcss v3.3.3 | MIT License | https://tailwindcss.com*//*1. Prevent padding and border from affecting element width. (https://github.com/mozdevs/cssremedy/issues/4)2. Allow adding a border to an element by just adding a border-width. (https://github.com/tailwindcss/tailwindcss/pull/116)*/*,::before,::after { box-sizing: border-box; border-width: 0; border-style: solid; border-color: #e5e7eb;}::before,::after { --tw-content: '';}/*1. Use a consistent sensible line-height in all browsers.2. Prevent adjustments of font size after orientation changes in iOS.3. Use a more readable tab size.4. Use the user's configured sans font-family by default.5. Use the user's configured sans font-feature-settings by default.6. Use the user's configured sans font-variation-settings by default.*/html { line-height: 1.5; -webkit-text-size-adjust: 100%; -moz-tab-size: 4; -o-tab-size: 4; tab-size: 4; font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"; font-feature-settings: normal; font-variation-settings: normal;}/*1. Remove the margin in all browsers.2. Inherit line-height from html so users can set them as a class directly on the html element.*/body { margin: 0; line-height: inherit;}/*1. Add the correct height in Firefox.2. Correct the inheritance of border color in Firefox. (https://bugzilla.mozilla.org/show_bug.cgi?id=190655)3. Ensure horizontal rules are visible by default.*/hr { height: 0; color: inherit; border-top-width: 1px;}/*Add the correct text decoration in Chrome, Edge, and Safari.*/abbr:where([title]) { -webkit-text-decoration: underline dotted; text-decoration: underline dotted;}/*Remove the default font size and weight for headings.*/h1,h2,h3,h4,h5,h6 { font-size: inherit; font-weight: inherit;}/*Reset links to optimize for opt-in styling instead of opt-out.*/a { color: inherit; text-decoration: inherit;}/*Add the correct font weight in Edge and Safari.*/b,strong { font-weight: bolder;}/*1. Use the user's configured mono font family by default.2. Correct the odd em font sizing in all browsers.*/code,kbd,samp,pre { font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace; font-size: 1em;}/*Add the correct font size in all browsers.*/small { font-size: 80%;}/*Prevent sub and sup elements from affecting the line height in all browsers.*/sub,sup { font-size: 75%; line-height: 0; position: relative; vertical-align: baseline;}sub { bottom: -0.25em;}sup { top: -0.5em;}/*1. Remove text indentation from table contents in Chrome and Safari. (https://bugs.chromium.org/p/chromium/issues/detail?id=999088, https://bugs.webkit.org/show_bug.cgi?id=201297)2. Correct table border color inheritance in all Chrome and Safari. (https://bugs.chromium.org/p/chromium/issues/detail?id=935729, https://bugs.webkit.org/show_bug.cgi?id=195016)3. Remove gaps between table borders by default.*/table { text-indent: 0; border-color: inherit; border-collapse: collapse;}/*1. Change the font styles in all browsers.2. Remove the margin in Firefox and Safari.3. Remove default padding in all browsers.*/button,input,optgroup,select,textarea { font-family: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 100%; font-weight: inherit; line-height: inherit; color: inherit; margin: 0; padding: 0;}/*Remove the inheritance of text transform in Edge and Firefox.*/button,select { text-transform: none;}/*1. Correct the inability to style clickable types in iOS and Safari.2. Remove default button styles.*/button,[type='button'],[type='reset'],[type='submit'] { -webkit-appearance: button; background-color: transparent; background-image: none;}/*Use the modern Firefox focus style for all focusable elements.*/:-moz-focusring { outline: auto;}/*Remove the additional :invalid styles in Firefox. (https://github.com/mozilla/gecko-dev/blob/2f9eacd9d3d995c937b4251a5557d95d494c9be1/layout/style/res/forms.css#L728-L737)*/:-moz-ui-invalid { box-shadow: none;}/*Add the correct vertical alignment in Chrome and Firefox.*/progress { vertical-align: baseline;}/*Correct the cursor style of increment and decrement buttons in Safari.*/::-webkit-inner-spin-button,::-webkit-outer-spin-button { height: auto;}/*1. Correct the odd appearance in Chrome and Safari.2. Correct the outline style in Safari.*/[type='search'] { -webkit-appearance: textfield; outline-offset: -2px;}/*Remove the inner padding in Chrome and Safari on macOS.*/::-webkit-search-decoration { -webkit-appearance: none;}/*1. Correct the inability to style clickable types in iOS and Safari.2. Change font properties to inherit in Safari.*/::-webkit-file-upload-button { -webkit-appearance: button; font: inherit;}/*Add the correct display in Chrome and Safari.*/summary { display: list-item;}/*Removes the default spacing and border for appropriate elements.*/blockquote,dl,dd,h1,h2,h3,h4,h5,h6,hr,figure,p,pre { margin: 0;}fieldset { margin: 0; padding: 0;}legend { padding: 0;}ol,ul,menu { list-style: none; margin: 0; padding: 0;}/*Reset default styling for dialogs.*/dialog { padding: 0;}/*Prevent resizing textareas horizontally by default.*/textarea { resize: vertical;}/*1. Reset the default placeholder opacity in Firefox. (https://github.com/tailwindlabs/tailwindcss/issues/3300)2. Set the default placeholder color to the user's configured gray 400 color.*/input::-moz-placeholder, textarea::-moz-placeholder { opacity: 1; color: #9ca3af;}input::placeholder,textarea::placeholder { opacity: 1; color: #9ca3af;}/*Set the default cursor for buttons.*/button,[role="button"] { cursor: pointer;}/*Make sure disabled buttons don't get the pointer cursor.*/:disabled { cursor: default;}/*1. Make replaced elements display: block by default. (https://github.com/mozdevs/cssremedy/issues/14)2. Add vertical-align: middle to align replaced elements more sensibly by default. (https://github.com/jensimmons/cssremedy/issues/14#issuecomment-634934210) This can trigger a poorly considered lint error in some tools but is included by design.*/img,svg,video,canvas,audio,iframe,embed,object { display: block; vertical-align: middle;}/*Constrain images and videos to the parent width and preserve their intrinsic aspect ratio. (https://github.com/mozdevs/cssremedy/issues/14)*/img,video { max-width: 100%; height: auto;}[hidden] { display: none;}*, ::before, ::after { --tw-border-spacing-x: 0; --tw-border-spacing-y: 0; --tw-translate-x: 0; --tw-translate-y: 0; --tw-rotate: 0; --tw-skew-x: 0; --tw-skew-y: 0; --tw-scale-x: 1; --tw-scale-y: 1; --tw-pan-x: ; --tw-pan-y: ; --tw-pinch-zoom: ; --tw-scroll-snap-strictness: proximity; --tw-gradient-from-position: ; --tw-gradient-via-position: ; --tw-gradient-to-position: ; --tw-ordinal: ; --tw-slashed-zero: ; --tw-numeric-figure: ; --tw-numeric-spacing: ; --tw-numeric-fraction: ; --tw-ring-inset: ; --tw-ring-offset-width: 0px; --tw-ring-offset-color: #fff; --tw-ring-color: rgb(59 130 246 / 0.5); --tw-ring-offset-shadow: 0 0 #0000; --tw-ring-shadow: 0 0 #0000; --tw-shadow: 0 0 #0000; --tw-shadow-colored: 0 0 #0000; --tw-blur: ; --tw-brightness: ; --tw-contrast: ; --tw-grayscale: ; --tw-hue-rotate: ; --tw-invert: ; --tw-saturate: ; --tw-sepia: ; --tw-drop-shadow: ; --tw-backdrop-blur: ; --tw-backdrop-brightness: ; --tw-backdrop-contrast: ; --tw-backdrop-grayscale: ; --tw-backdrop-hue-rotate: ; --tw-backdrop-invert: ; --tw-backdrop-opacity: ; --tw-backdrop-saturate: ; --tw-backdrop-sepia: ;}::backdrop { --tw-border-spacing-x: 0; --tw-border-spacing-y: 0; --tw-translate-x: 0; --tw-translate-y: 0; --tw-rotate: 0; --tw-skew-x: 0; --tw-skew-y: 0; --tw-scale-x: 1; --tw-scale-y: 1; --tw-pan-x: ; --tw-pan-y: ; --tw-pinch-zoom: ; --tw-scroll-snap-strictness: proximity; --tw-gradient-from-position: ; --tw-gradient-via-position: ; --tw-gradient-to-position: ; --tw-ordinal: ; --tw-slashed-zero: ; --tw-numeric-figure: ; --tw-numeric-spacing: ; --tw-numeric-fraction: ; --tw-ring-inset: ; --tw-ring-offset-width: 0px; --tw-ring-offset-color: #fff; --tw-ring-color: rgb(59 130 246 / 0.5); --tw-ring-offset-shadow: 0 0 #0000; --tw-ring-shadow: 0 0 #0000; --tw-shadow: 0 0 #0000; --tw-shadow-colored: 0 0 #0000; --tw-blur: ; --tw-brightness: ; --tw-contrast: ; --tw-grayscale: ; --tw-hue-rotate: ; --tw-invert: ; --tw-saturate: ; --tw-sepia: ; --tw-drop-shadow: ; --tw-backdrop-blur: ; --tw-backdrop-brightness: ; --tw-backdrop-contrast: ; --tw-backdrop-grayscale: ; --tw-backdrop-hue-rotate: ; --tw-backdrop-invert: ; --tw-backdrop-opacity: ; --tw-backdrop-saturate: ; --tw-backdrop-sepia: ;}.container { width: 100%;}@media (min-width: 640px) { .container { max-width: 640px; }}@media (min-width: 768px) { .container { max-width: 768px; }}@media (min-width: 1024px) { .container { max-width: 1024px; }}@media (min-width: 1280px) { .container { max-width: 1280px; }}@media (min-width: 1536px) { .container { max-width: 1536px; }}.visible { visibility: visible;}.\!visible { visibility: visible !important;}.fixed { position: fixed;}.relative { position: relative;}.bottom-6 { bottom: 1.5rem;}.right-4 { right: 1rem;}.mx-auto { margin-left: auto; margin-right: auto;}.my-8 { margin-top: 2rem; margin-bottom: 2rem;}.my-4 { margin-top: 1rem; margin-bottom: 1rem;}.mb-3 { margin-bottom: 0.75rem;}.mb-4 { margin-bottom: 1rem;}.ml-1 { margin-left: 0.25rem;}.mt-\[30vh\] { margin-top: 30vh;}.block { display: block;}.inline { display: inline;}.flex { display: flex;}.hidden { display: none;}.h-8 { height: 2rem;}.h-fit { height: -moz-fit-content; height: fit-content;}.h-full { height: 100%;}.h-screen { height: 100vh;}.min-h-screen { min-height: 100vh;}.w-11\/12 { width: 91.666667%;}.w-12 { width: 3rem;}.w-8 { width: 2rem;}.w-8\/12 { width: 66.666667%;}.w-full { width: 100%;}.w-2\/5 { width: 40%;}.flex-col { flex-direction: column;}.items-center { align-items: center;}.justify-center { justify-content: center;}.gap-2 { gap: 0.5rem;}.gap-6 { gap: 1.5rem;}.overflow-hidden { overflow: hidden;}.overflow-x-scroll { overflow-x: scroll;}.scroll-smooth { scroll-behavior: smooth;}.text-ellipsis { text-overflow: ellipsis;}.whitespace-nowrap { white-space: nowrap;}.rounded { border-radius: 0.25rem;}.rounded-full { border-radius: 9999px;}.border { border-width: 1px;}.border-slate-100 { --tw-border-opacity: 1; border-color: rgb(241 245 249 / var(--tw-border-opacity));}.bg-gray-100 { --tw-bg-opacity: 1; background-color: rgb(243 244 246 / var(--tw-bg-opacity));}.bg-gray-50\/40 { background-color: rgb(249 250 251 / 0.4);}.bg-inherit { background-color: inherit;}.fill-black { fill: #000;}.fill-blue-600 { fill: #2563eb;}.fill-inherit { fill: inherit;}.p-1 { padding: 0.25rem;}.px-2 { padding-left: 0.5rem; padding-right: 0.5rem;}.py-1 { padding-top: 0.25rem; padding-bottom: 0.25rem;}.py-2 { padding-top: 0.5rem; padding-bottom: 0.5rem;}.py-4 { padding-top: 1rem; padding-bottom: 1rem;}.px-4 { padding-left: 1rem; padding-right: 1rem;}.px-6 { padding-left: 1.5rem; padding-right: 1.5rem;}.py-3 { padding-top: 0.75rem; padding-bottom: 0.75rem;}.px-1 { padding-left: 0.25rem; padding-right: 0.25rem;}.pb-16 { padding-bottom: 4rem;}.pt-4 { padding-top: 1rem;}.text-center { text-align: center;}.text-2xl { font-size: 1.5rem; line-height: 2rem;}.text-lg { font-size: 1.125rem; line-height: 1.75rem;}.text-xl { font-size: 1.25rem; line-height: 1.75rem;}.font-bold { font-weight: 700;}.shadow-2xl { --tw-shadow: 0 25px 50px -12px rgb(0 0 0 / 0.25); --tw-shadow-colored: 0 25px 50px -12px var(--tw-shadow-color); box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);}.shadow-lg { --tw-shadow: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1); --tw-shadow-colored: 0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color); box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);}.duration-300 { transition-duration: 300ms;}.ease-in { transition-timing-function: cubic-bezier(0.4, 0, 1, 1);}.hover\:underline:hover { text-decoration-line: underline;}:is(.dark .dark\:border-slate-600) { --tw-border-opacity: 1; border-color: rgb(71 85 105 / var(--tw-border-opacity));}:is(.dark .dark\:border-slate-500) { --tw-border-opacity: 1; border-color: rgb(100 116 139 / var(--tw-border-opacity));}:is(.dark .dark\:border-gray-700) { --tw-border-opacity: 1; border-color: rgb(55 65 81 / var(--tw-border-opacity));}:is(.dark .dark\:bg-black) { --tw-bg-opacity: 1; background-color: rgb(0 0 0 / var(--tw-bg-opacity));}:is(.dark .dark\:bg-gray-900\/30) { background-color: rgb(17 24 39 / 0.3);}:is(.dark .dark\:bg-gray-900\/90) { background-color: rgb(17 24 39 / 0.9);}:is(.dark .dark\:fill-blue-300) { fill: #93c5fd;}:is(.dark .dark\:fill-slate-50) { fill: #f8fafc;}:is(.dark .dark\:text-gray-50) { --tw-text-opacity: 1; color: rgb(249 250 251 / var(--tw-text-opacity));}:is(.dark .dark\:shadow-gray-700\/40) { --tw-shadow-color: rgb(55 65 81 / 0.4); --tw-shadow: var(--tw-shadow-colored);}:is(.dark .dark\:hover\:bg-gray-800:hover) { --tw-bg-opacity: 1; background-color: rgb(31 41 55 / var(--tw-bg-opacity));}@media (min-width: 768px) { .md\:w-3\/4 { width: 75%; } .md\:w-28 { width: 7rem; } .md\:w-1\/4 { width: 25%; } .md\:w-2\/5 { width: 40%; } .md\:py-2 { padding-top: 0.5rem; padding-bottom: 0.5rem; } .md\:text-xl { font-size: 1.25rem; line-height: 1.75rem; } .md\:shadow-none { --tw-shadow: 0 0 #0000; --tw-shadow-colored: 0 0 #0000; box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow); } .md\:hover\:shadow-lg:hover { --tw-shadow: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1); --tw-shadow-colored: 0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color); box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow); } .md\:hover\:ease-in-out:hover { transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1); }}
    </style>
    <link rel="icon" href="/dQw4w9WgXcQ/favicon.svg" />
  </head>
  <body class="dark relative">
    <div class="dark:bg-black dark:text-gray-50 ease-in duration-300 min-h-screen">
      <!-- container -->
      <div class="w-11/12 md:w-3/4 mx-auto pb-16 pt-4">
        <!-- navigation -->
        <input type="text" id="search" class="border dark:border-gray-700 bg-inherit rounded-full px-4 py-1 mx-auto block text-xl md:w-2/5 my-4" placeholder="Search"/>
        <div class="flex items-center py-2 px-2 mb-4 overflow-x-scroll scroll-smooth text-ellipsis whitespace-nowrap">
          {{range .ProgessPah}}
          <a href="{{.Url}}" class="font-bold hover:underline text-lg"> {{if .SlashPre}}/{{end}}{{.Title}}</a>
          {{end}}
        </div>

        <!-- diretoris and files list -->
        <div class="my-8">
          {{range .Directories }}
          <a fill-black dark:fill-slate-50 href="{{.Url}}" class="files flex gap-6 items-center px-2 py-1 md:py-2 mb-3 bg--100 bg-gray-100 dark:bg-gray-900/90 shadow-lg dark:hover:bg-gray-800 dark:shadow-gray-700/40 md:shadow-none md:hover:shadow-lg duration-300 md:hover:ease-in-out">
            <div class="ml-1 w-8/12 flex items-center gap-2">
              <span>{{ .Icon }}</span>
              <span class="name overflow-hidden text-ellipsis w-full whitespace-nowrap">{{.Name}}</span >
            </div>
            <div class="overflow-hidden text-ellipsis whitespace-nowrap">
              {{if .IsDir}} -- {{else}}{{.Size}}{{end}}
            </div>
          </a>
          {{else}}
          <div class="text-center font-bold w-full h-screen flex justify-center mt-[30vh] text-2xl">
            <span>No files</span>
          </div>
          {{end}}
        </div>

        <!-- navigation -->
        <div class="fixed right-4 bottom-6 w-12 h-fit flex flex-col items-center justify-center gap-2">
          <!-- menu items -->
          <div id="menu-items" class="hidden bg-gray-50/40 dark:bg-gray-900/30 px-2 py-2 rounded border border-slate-100 dark:border-slate-600 shadow-2xl">
            <div class="flex gap-2 flex-col items-center justify-center fill-blue-600 dark:fill-blue-300">
              <button id="theme-btn" class="w-8 h-8"></button>
              <button id="menu-cancel" class="w-8 h-8">
                <svg class="w-full h-full fill-inherit" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 16 16" > <path d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z" /> </svg>
              </button>
            </div>
          </div>

          <!-- menu tougle -->
          <div id="menu" class="w-8 h-8 p-1">
            <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" class="bi bi-three-dots w-full h-full fill-blue-600 dark:fill-blue-300" viewBox="0 0 16 16" > <path d="M3 9.5a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3zm5 0a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3zm5 0a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3z" /> </svg>
          </div>

          <a href="{{.PreviousPage}}" class="w-8 h-8 block">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="w-full h-full fill-blue-600 dark:fill-blue-300" viewBox="0 0 16 16" > <path fill-rule="evenodd"
                d="M12 8a.5.5 0 0 1-.5.5H5.707l2.147 2.146a.5.5 0 0 1-.708.708l-3-3a.5.5 0 0 1 0-.708l3-3a.5.5 0 1 1 .708.708L5.707 7.5H11.5a.5.5 0 0 1 .5.5z"
              />
            </svg>
          </a>
        </div>
      </div>
    </div>
    <script>
      const light = '<svg class="w-full h-full fill-inherit" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><title>brightness-6</title><path d="M12,18V6A6,6 0 0,1 18,12A6,6 0 0,1 12,18M20,15.31L23.31,12L20,8.69V4H15.31L12,0.69L8.69,4H4V8.69L0.69,12L4,15.31V20H8.69L12,23.31L15.31,20H20V15.31Z" /></svg>';
      const dark = '<svg class="w-full h-full fill-inherit" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><title>brightness-4</title><path d="M12,18C11.11,18 10.26,17.8 9.5,17.45C11.56,16.5 13,14.42 13,12C13,9.58 11.56,7.5 9.5,6.55C10.26,6.2 11.11,6 12,6A6,6 0 0,1 18,12A6,6 0 0,1 12,18M20,8.69V4H15.31L12,0.69L8.69,4H4V8.69L0.69,12L4,15.31V20H8.69L12,23.31L15.31,20H20V15.31L23.31,12L20,8.69Z" /></svg>';

      const body = document.body;
      const theme_btn = document.getElementById("theme-btn");
      const menu = document.getElementById("menu");
      const menu_items = document.getElementById("menu-items");
      const menu_cancel = document.getElementById("menu-cancel");

      menu.addEventListener("click", () => {
        menu.style.display = "none";
        menu_items.classList.remove("hidden");
      });

      menu_cancel.addEventListener("click", () => {
        menu.style.display = "block";
        menu_items.classList.add("hidden");
      });

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

      const files = document.querySelectorAll(".files");
      document.getElementById("search").addEventListener("input", (e) => {
        let value = e.target.value;
        files.forEach((file) => {
          const visible = file
            .querySelector(".name")
            .innerText.toLowerCase()
            .includes(value);

          file.classList.toggle("hidden", !visible);
        });
      });
    </script>
  </body>
</html>
{{end}}`
