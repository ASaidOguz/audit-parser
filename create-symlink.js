// create-symlink.js
const { execSync } = require('child_process');
const os = require('os');

const platform = os.platform();

if (platform === 'win32') {
  // Create symbolic link for Windows using mklink command
  execSync('mklink .\\bin\\audit-parser .\\node_modules\\.bin\\audit-parser', { stdio: 'inherit' });
} else if (platform === 'darwin' || platform === 'linux') {
  // Create symbolic link for Unix-like systems using ln command
  execSync('ln -sf ./node_modules/.bin/audit-parser ./bin/audit-parser', { stdio: 'inherit' });
} else {
  console.error('Unsupported operating system');
}
