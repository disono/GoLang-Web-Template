let mix = require('laravel-mix');

// l = local, p = production
let env = 'l';

/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel application. By default, we are compiling the Sass
 | file for the application as well as bundling up all the JS files.
 |
 */

let vendorJS = 'resources/static/public/js/vendor.js';
let vendorCSS = 'resources/static/public/css/vendor.css';

let WB_JS = [
    'node_modules/uikit/dist/js/uikit.min.js',
    'node_modules/uikit/dist/js/uikit-icons.min.js'
];

let WB_CSS = [
    'node_modules/uikit/dist/css/uikit.min.css'
];

if (env === 'p') {
    console.log('Processing Production...');

    mix.scripts(WB_JS.concat([
        'resources/static/js/app.js'
    ]), vendorJS);

    mix.styles(WB_CSS.concat([
        'public/assets/css/theme.css'
    ]), vendorCSS);
} else if (env === 'l') {
    console.log('Processing Development...');

    mix.scripts(WB_JS, vendorJS);
    mix.styles(WB_CSS, vendorCSS);
}