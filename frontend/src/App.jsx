import { useState, useEffect } from 'react';

// Import components from the components-lib submodule
import { Button, Card, Alert, Input, Spinner } from '../components-lib/src/index.js';

// Import utilities from the utils-lib submodule
import { 
  formatDate, 
  getTimeAgo, 
  capitalize, 
  truncate, 
  isEmail,
  formatCurrency,
  getLibraryInfo
} from '../utils-lib/src/index.js';

function App() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [emailError, setEmailError] = useState('');
  const [apiData, setApiData] = useState(null);
  const [loading, setLoading] = useState(false);
  const [showAlert, setShowAlert] = useState(true);
  const [greetingResponse, setGreetingResponse] = useState(null);

  // Validate email using utils-lib
  const handleEmailChange = (e) => {
    const value = e.target.value;
    setEmail(value);
    if (value && !isEmail(value)) {
      setEmailError('Please enter a valid email address');
    } else {
      setEmailError('');
    }
  };

  // Fetch data from backend API
  const fetchAPIData = async () => {
    setLoading(true);
    try {
      const response = await fetch('/api/');
      const data = await response.json();
      setApiData(data);
    } catch (error) {
      console.error('Error fetching API data:', error);
    } finally {
      setLoading(false);
    }
  };

  // Fetch greeting from backend
  const fetchGreeting = async () => {
    if (!name) return;
    setLoading(true);
    try {
      const response = await fetch(`/api/greet?name=${encodeURIComponent(name)}`);
      const data = await response.json();
      setGreetingResponse(data);
    } catch (error) {
      console.error('Error fetching greeting:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchAPIData();
  }, []);

  const currentDate = new Date();

  return (
    <div>
      <header className="header">
        <h1>🚀 Git Submodule POC - Frontend</h1>
        <p>Demonstrating Modular Frontend with React and Git Submodules</p>
        <div style={{ marginTop: '10px' }}>
          <span className="info-badge">Components-Lib Submodule</span>
          <span className="info-badge">Utils-Lib Submodule</span>
          <span className="info-badge">React + Vite</span>
        </div>
      </header>

      <div className="container">
        {/* Alert Demo */}
        {showAlert && (
          <Alert
            type="success"
            message="🎉 Frontend successfully loaded with two Git submodule libraries!"
            onClose={() => setShowAlert(false)}
          />
        )}

        {/* Backend Integration Section */}
        <section className="section">
          <h2>📡 Backend API Integration</h2>
          <p style={{ marginBottom: '15px' }}>
            This frontend connects to the Go Echo backend running on port 8080
          </p>
          
          {loading && (
            <div style={{ display: 'flex', justifyContent: 'center', padding: '20px' }}>
              <Spinner size="medium" />
            </div>
          )}

          {apiData && !loading && (
            <div className="api-result">
              <strong>Backend Response:</strong>
              <pre>{JSON.stringify(apiData, null, 2)}</pre>
            </div>
          )}

          <div style={{ marginTop: '20px' }}>
            <Button variant="primary" onClick={fetchAPIData} disabled={loading}>
              🔄 Refresh Backend Data
            </Button>
          </div>
        </section>

        {/* Components Library Demo */}
        <section className="section">
          <h2>🎨 Components Library (Submodule 1)</h2>
          <p style={{ marginBottom: '20px' }}>
            These components are loaded from <code>frontend/components-lib/</code>
          </p>

          <div className="grid">
            <Card 
              title="Greeting Form" 
              footer={greetingResponse && (
                <small>
                  Last greeting: {getTimeAgo(new Date())}
                </small>
              )}
            >
              <Input
                label="Your Name"
                placeholder="Enter your name"
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
              <Input
                label="Email Address"
                type="email"
                placeholder="Enter your email"
                value={email}
                onChange={handleEmailChange}
                error={emailError}
              />
              <Button 
                variant="success" 
                onClick={fetchGreeting}
                disabled={!name || loading}
              >
                Send Greeting to Backend
              </Button>

              {greetingResponse && (
                <div className="api-result" style={{ marginTop: '15px' }}>
                  <strong>Greeting Response:</strong>
                  <pre>{JSON.stringify(greetingResponse, null, 2)}</pre>
                </div>
              )}
            </Card>

            <Card title="Button Variants">
              <div style={{ display: 'flex', flexDirection: 'column', gap: '10px' }}>
                <Button variant="primary">Primary Button</Button>
                <Button variant="secondary">Secondary Button</Button>
                <Button variant="success">Success Button</Button>
                <Button variant="danger">Danger Button</Button>
                <Button variant="primary" disabled>Disabled Button</Button>
              </div>
            </Card>
          </div>
        </section>

        {/* Utilities Library Demo */}
        <section className="section">
          <h2>🛠️ Utilities Library (Submodule 2)</h2>
          <p style={{ marginBottom: '20px' }}>
            These utilities are loaded from <code>frontend/utils-lib/</code>
          </p>

          <div className="grid">
            <Card title="Date Formatting">
              <p><strong>Current Date:</strong></p>
              <p>YYYY-MM-DD: {formatDate(currentDate, 'YYYY-MM-DD')}</p>
              <p>DD/MM/YYYY: {formatDate(currentDate, 'DD/MM/YYYY')}</p>
              <p>Time Ago: {getTimeAgo(new Date(Date.now() - 3600000))}</p>
            </Card>

            <Card title="String Utilities">
              <p><strong>capitalize:</strong> {capitalize('hello world')}</p>
              <p><strong>truncate:</strong> {truncate('This is a very long text that will be truncated', 20)}</p>
            </Card>

            <Card title="Number Formatting">
              <p><strong>Currency:</strong> {formatCurrency(1234.56, 'USD')}</p>
              <p><strong>Currency (EUR):</strong> {formatCurrency(9876.54, 'EUR', 'en-US')}</p>
            </Card>

            <Card title="Library Info">
              <p>{getLibraryInfo()}</p>
              <p style={{ marginTop: '10px', fontSize: '14px', color: '#666' }}>
                This demonstrates that the utilities library is successfully loaded as a Git submodule.
              </p>
            </Card>
          </div>
        </section>

        {/* Architecture Info */}
        <section className="section">
          <h2>🏗️ Architecture</h2>
          <Card>
            <h3 style={{ marginBottom: '15px', fontSize: '1.2rem' }}>Frontend Modular Structure</h3>
            <pre style={{ 
              backgroundColor: '#f8f9fa', 
              padding: '15px', 
              borderRadius: '4px',
              fontSize: '13px',
              overflow: 'auto'
            }}>
{`frontend/
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
│   ├── main.jsx
│   └── index.css
├── package.json
└── vite.config.js`}
            </pre>

            <div style={{ marginTop: '20px' }}>
              <h4 style={{ marginBottom: '10px' }}>Benefits:</h4>
              <ul style={{ paddingLeft: '20px', lineHeight: '1.8' }}>
                <li>✅ Modular and reusable components</li>
                <li>✅ Independent versioning of libraries</li>
                <li>✅ Easy to share across multiple projects</li>
                <li>✅ Clean separation of concerns</li>
                <li>✅ Git submodules for dependency management</li>
              </ul>
            </div>
          </Card>
        </section>
      </div>
    </div>
  );
}

export default App;
