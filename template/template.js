function getHtml() {
    return `
    <!DOCTYPE html>
        <html lang="en">
            <head>
                <meta charset="UTF-8">
                <title>Title</title>
                <link rel="preload" href="http://localhost:3001/assets/fonts/manrope/manrope-medium.woff2" as="font" type="font/woff2">
                <link rel="preload" href="http://localhost:3001/assets/fonts/manrope/manrope-bold.woff2" as="font" type="font/woff2">
                <link rel="preload" href="http://localhost:3001/assets/fonts/arial-rub/arial-rub.woff2" as="font" type="font/woff2">
                <link rel="preload" href="http://localhost:3001/assets/fonts/arial-rub/arial-rub-bold.woff2" as="font" type="font/woff2">
                
                <style>
                    *, *::after, *::before {
                        animation-delay: -0.0001s !important;
                        animation-duration: 0s !important;
                        animation-play-state: paused !important;
                        transition-delay: 0s !important;
                        transition-duration: 0s !important;
                        caret-color: transparent !important;
                    }

                    @font-face {
                        font-family: Manrope;
                        src: url('http://localhost:3001/assets/fonts/manrope/manrope-medium.woff2') format('woff2');
                        font-weight: 100 400;
                    }

                    @font-face {
                        font-family: Manrope;
                        src: url('http://localhost:3001/assets/fonts/manrope/manrope-bold.woff2') format('woff2');
                        font-weight: 700 800;
                    }

                    @font-face {
                        font-family: 'Arial Rub';
                        font-weight: normal;
                        src: url('http://localhost:3001/assets/fonts/arial-rub/arial-rub.woff2') format('woff2');
                    }

                    @font-face {
                        font-family: 'Arial Rub';
                        font-weight: bold;
                        src: url('http://localhost:3001/assets/fonts/arial-rub/arial-rub-bold.woff2') format('woff2');
                    }

                    .container {
                        width: 100%;
                    }

                    .button {
                        width: 200px;
                        height: 60px;
                        background: rebeccapurple;
                        color: white;
                        border-radius: 10px;
                    }

                    .test {
                        font-family: "Arial Rub";
                        font-weight: 800;
                        font-size: 20px;
                    }
                </style>
            </head>
            <body>
                <div class="contaner">
                    <button class="button">
                        <span class="test">Hellow world</span>
                    </button>
                </div>
            </body>
        </html>
`;
}

function generateHtml(raw = false) {
    const html = getHtml();
    if(raw) return html
    const arrayStrings = html.split('\n').map((item) => {
        return item.trim()
    })
    return arrayStrings.join('')
}

module.exports = generateHtml