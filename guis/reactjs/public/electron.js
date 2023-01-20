const { app, BrowserWindow } = require('electron');
const isDev = require('electron-is-dev');

function createWindow() {
    // Create the browser window.
    const win = new BrowserWindow({
        width: 2048,
        height: 1456,
        webPreferences: {
            nodeIntegration: true,
        },
    });

    // if we are is dev mode:
    if (isDev) {
        // load ur:
        win.loadURL('http://localhost:3000');

        // Open the DevTools.
        win.webContents.openDevTools({ mode: 'attached' });

        // add the react extension to the dev tools:
        const { default: installExtension, REACT_DEVELOPER_TOOLS } = require('electron-devtools-installer');
        const { app } = require('electron');
        app.whenReady().then(() => {
        installExtension(REACT_DEVELOPER_TOOLS)
            .then((name) => console.log(`Added Extension:  ${name}`))
            .catch((err) => console.log('An error occurred: ', err));
        });

        return;
    };

    // load url:
    win.loadURL(`file://${path.join(__dirname, '../build/index.html')}`);
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.whenReady().then(createWindow);

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit();
    };
});

app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) {
        createWindow();
    };
});
