import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import BookNow from './components/BookNow';
import Header from './components/header';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Header />} />
        <Route path="/booknow" element={<BookNow />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
