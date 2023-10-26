import './App.css';
import  Footer  from './Components/FooterComponent/Footer.js';
import Home from './Components/HomeComponent/Home';
import {  Routes, Route } from 'react-router-dom';
import  Booking  from './Components/BookingComponent/Booking.js';
import  About  from './Components/AboutComponents/About';
import  Contact  from './Components/ContactComponents/Contact';
import Header from './Components/HeaderComponent/Header';
function App() {

  return (
    
    <div >
      <Routes>
        <Route path="/" element={<><Home />
      <Footer/></>} />

        <Route
          path="/booking"
          element={
            <>
              <Header />
              <Booking />
              <Footer />
            </>
          }
        />
        <Route
          path="/about"
          element={
            <>
              <Header />
              <About />
              <Footer />
            </>
          }
        />
        <Route
          path="/contact"
          element={
            <>
              <Header />
              <Contact />
              <Footer />
            </>
          }
        />
      </Routes>

      </div>
  );
}

export default App;
