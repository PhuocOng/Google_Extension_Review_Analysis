import React, { useState } from 'react';
import axios from 'axios';
import './App.css'; // Optional if you want to style it

function App() {
  const [url, setUrl] = useState(''); // To store the URL input
  const [result, setResult] = useState(null); // To store the result from the backend
  const [error, setError] = useState(null); // To store errors, if any

  const handleSubmit = async (e) => {
    e.preventDefault(); // Prevent page reload on form submit

    try {
      // Send POST request to the backend API
      const response = await axios.post('http://127.0.0.1:5000/api/analyze-extension', { url });
      setResult(response.data); // Store result in state
      setError(null); // Reset error if request is successful
    } catch (err) {
      console.error('Error:', err);
      setError('Failed to fetch the analysis. Please check the URL and try again.');
    }

    console.log(result)
  };


  return (
    <div className="App">
      <h1>Chrome Extension Review Analyzer</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          placeholder="Enter Chrome Extension URL"
          required
        />
        <button type="submit">Analyze</button>
      </form>

      {error && <p style={{ color: 'red' }}>{error}</p>}
      {result && (
        <div className="results">
          <h2>Analysis Results:</h2>

          {/* Strengths Section */}
          <h3>Strengths</h3>
          <p><strong>Summary:</strong> {result["strengths"]}</p>
          <ul>
            {result["strengths_arr"].map((strength, index) => (
              <li key={index}>{strength}</li>
            ))}
          </ul>

          {/* Weaknesses Section */}
          <h3>Weaknesses</h3>
          <p><strong>Summary:</strong> {result["weaknesses"]}</p>
          <ul>
            {result["weaknesses_arr"].map((weakness, index) => (
              <li key={index}>{weakness}</li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}

export default App;
