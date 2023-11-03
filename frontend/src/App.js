import React, { useEffect, useState } from 'react';
import './App.css';
import { Routes, Route } from 'react-router-dom';
import Footer from './components/Footer';
import Section2 from './components/Section2';
import Header from './components/header';
import Section1 from './components/section1';
import BookingForm from './components/BookNow';

function App() {
  const [Heading, setHeading] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/admin/heading/active')
      .then((response) => response.json())
      .then((data) => setHeading(data.data))
      .catch((error) => console.error('Error fetching Heading data:', error));
  }, []);

  return (
    <Routes>
      <Route path='/' element={
        <div className='body'>
          <Header />
          <Section1 Headingdata={Heading} />
          <Section2 Headingdata={Heading} />
          <Footer Headingdata={Heading} />
        </div>
      } />

      <Route path='/BookNow' element={<BookingForm />} />
    </Routes>
  );
}

export default App;
