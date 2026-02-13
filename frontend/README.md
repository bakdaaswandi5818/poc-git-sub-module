# Frontend Application - Git Submodule POC

A React-based frontend application demonstrating modular architecture with Git submodules.

## 🎯 Overview

This frontend application showcases how to build a modular React application using **Git submodules** for component and utility libraries. It connects to the Go Echo backend and demonstrates seamless integration between frontend modules and backend APIs.

## 📦 Submodule Libraries

### 1. Components Library (`components-lib/`)
A reusable React UI components library containing:
- **Button** - Multiple variants (primary, secondary, success, danger)
- **Card** - Container component with header and footer
- **Alert** - Notification messages with different types
- **Input** - Form input with label and error handling
- **Spinner** - Loading indicator

### 2. Utils Library (`utils-lib/`)
A comprehensive utility functions library containing:
- **Date Formatting** - formatDate, getTimeAgo
- **String Utilities** - capitalize, truncate, slugify
- **Validation** - isEmail, isURL, isPhoneNumber, validatePassword
- **Number Formatting** - formatCurrency, formatNumber
- **Array Utilities** - groupBy, sortBy, unique, chunk
- **Storage** - localStorage wrapper
- **Performance** - debounce, throttle
- **Deep Clone** - Object cloning utility

## 🏗️ Architecture

```
frontend/
├── components-lib/          # Git Submodule 1
│   ├── src/index.js        # React components
│   ├── package.json
│   └── README.md
├── utils-lib/              # Git Submodule 2
│   ├── src/index.js        # Utility functions
│   ├── package.json
│   └── README.md
├── src/
│   ├── App.jsx             # Main application
│   ├── main.jsx            # Entry point
│   └── index.css           # Global styles
├── package.json
├── vite.config.js
└── index.html
```

## 🚀 Getting Started

### Prerequisites
- Node.js 16+ and npm
- Git
- Backend server running on port 8080

### Installation

1. **Clone with submodules:**
```bash
git clone --recurse-submodules <repository-url>
cd poc-git-sub-module/frontend
```

2. **Or initialize submodules after cloning:**
```bash
git clone <repository-url>
cd poc-git-sub-module
git submodule update --init --recursive
cd frontend
```

3. **Install dependencies:**
```bash
npm install
```

### Running the Application

```bash
# Development server (with hot reload)
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

The application will be available at `http://localhost:3000`

## 🔗 Backend Integration

The frontend connects to the Go Echo backend through a proxy configuration in `vite.config.js`:

```javascript
proxy: {
  '/api': {
    target: 'http://localhost:8080',
    changeOrigin: true,
    rewrite: (path) => path.replace(/^\/api/, '')
  }
}
```

**API Endpoints Used:**
- `GET /api/` - Get backend info with submodule versions
- `GET /api/greet?name=<name>` - Get personalized greeting
- `GET /api/health` - Health check

## 📚 Using Submodule Libraries

### Importing Components
```javascript
import { Button, Card, Alert, Input, Spinner } from '../components-lib/src/index.js';

// Usage
<Button variant="primary" onClick={handleClick}>
  Click Me
</Button>

<Card title="My Card">
  <p>Card content</p>
</Card>
```

### Importing Utilities
```javascript
import { 
  formatDate, 
  isEmail, 
  formatCurrency,
  storage 
} from '../utils-lib/src/index.js';

// Usage
const formatted = formatDate(new Date(), 'YYYY-MM-DD');
const isValid = isEmail('test@example.com');
const price = formatCurrency(1234.56, 'USD');
```

## 🔄 Working with Submodules

### Update submodules to latest:
```bash
git submodule update --remote --merge
```

### Make changes in a submodule:
```bash
cd components-lib
# Make changes
git add .
git commit -m "Update components"
git push origin master

cd ..
git add components-lib
git commit -m "Update components-lib reference"
```

### Check submodule status:
```bash
git submodule status
```

## 🎨 Features Demonstrated

1. **Modular Component Library** - Reusable React components as a separate submodule
2. **Utility Functions Library** - Shared utilities across projects
3. **Backend Integration** - Seamless communication with Go Echo API
4. **Real-time Data** - Fetching and displaying backend data
5. **Form Validation** - Using utility library for email validation
6. **Responsive Design** - Mobile-friendly layout
7. **Modern Build Setup** - Vite for fast development and production builds

## 🛠️ Technology Stack

- **React 18** - UI library
- **Vite** - Build tool and dev server
- **Git Submodules** - For modular architecture
- **CSS3** - Styling
- **Fetch API** - Backend communication

## 📖 Framework Recommendation

**Why React works well with Git Submodules:**

1. ✅ **ES Modules Support** - Easy to import from submodules
2. ✅ **Component-Based** - Natural fit for modular architecture
3. ✅ **Build Tool Flexibility** - Vite, Webpack, etc. handle relative imports
4. ✅ **No Framework Lock-in** - Submodules can be pure JavaScript
5. ✅ **Large Ecosystem** - Well-established patterns and tooling

**Other frameworks that work well:**
- Vue.js (similar module system)
- Svelte (compile-time optimization)
- Angular (with proper module configuration)
- Plain JavaScript/TypeScript

## 🎯 Benefits of This Architecture

1. **Reusability** - Components and utils can be used in multiple projects
2. **Independent Versioning** - Each submodule has its own version
3. **Team Collaboration** - Different teams can work on different modules
4. **Clean Separation** - Clear boundaries between modules
5. **Easy Updates** - Pull latest changes from submodule repos
6. **Consistent UI** - Shared component library ensures consistency

## 🔍 Development Tips

1. **Always initialize submodules** after cloning
2. **Commit submodule changes separately** before updating parent repo
3. **Use relative imports** for submodule code
4. **Document submodule dependencies** clearly
5. **Version your submodules** for stability

## 📝 License

MIT

## 🤝 Contributing

When contributing to submodules:
1. Make changes in the submodule directory
2. Commit and push submodule changes
3. Update parent repository reference
4. Test integration before pushing

## 🐛 Troubleshooting

### Submodules not loaded
```bash
git submodule update --init --recursive
```

### Cannot find module errors
```bash
# Check if submodules are initialized
ls -la components-lib/
ls -la utils-lib/

# If empty, initialize them
git submodule update --init --recursive
```

### Backend connection issues
- Ensure backend is running on port 8080
- Check proxy configuration in vite.config.js
- Verify CORS settings on backend

## 📚 Additional Resources

- [React Documentation](https://react.dev/)
- [Vite Documentation](https://vitejs.dev/)
- [Git Submodules Guide](../SUBMODULE_PULL_GUIDE.md)
- [Main Project README](../README.md)
