import React, { useEffect, useState } from "react";
import "./App.css";
import { Routes, Route } from "react-router-dom";
import BookingForm from "./components/BookNow";
import Home from "./components/Home";


function App() {
  
  return (

    <Routes>
      <Route path="/" element={<Home/>}/>
      <Route path="/BookNow" element={<BookingForm />} />      
    </Routes>

  );
}

export default App;
