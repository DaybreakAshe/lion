import '../../src/assets/css/App.css';
import Home from '../layout/home/Home.tsx';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/*" element={<Home />} />
        </Routes>
      </Router>
    </>
  );
}
export default App;
