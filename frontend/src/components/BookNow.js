import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import "./BookNow.css";
import availableSlot from "./Booking/AvailableSlot"

function BookingForm() {
  const [selectedDate, setSelectedDate] = useState(null);
  const [dropdownOptions, setDropdownOptions] = useState([]);
  const [selectedOption, setSelectedOption] = useState("");
  const [availableSlots, setAvailableSlots] = useState([]);

  const Postdata = () => {
    fetch("http://127.0.0.1:8080/user/get/avl/slots", {
      method: "POST",
      
      body: JSON.stringify({
        "branch_id": parseInt(selectedOption),
        "date": formatDate(selectedDate)
        
      }),
    })
      .then((res) => {
        console.log(res.data[0].Is_booked);
        if (!res.ok) {
          throw new Error("Network response was not ok");
        }
        return res.json();
      })
      .then((data) => {
        console.log(data.Slots);
        setAvailableSlots(data.available_slots);
      })
      .catch((error) => {
        console.error("Error while fetching data:", error);
      });
  };

  useEffect(() => {
    fetch("http://localhost:8080/admin/active/branch")
      .then((res) => res.json())
      .then((data) => {
        setDropdownOptions(data.data);
      });
  }, []);

  const handleSubmit = (e) => {
    e.preventDefault();

    console.log("Booking Date (dd/mm/yyyy):", formatDate(selectedDate));
    console.log("Selected Option:", selectedOption);

    // Reset text fields and close the form
    setSelectedDate(null);
    setSelectedOption("");
  };

  const formatDate = (date) => {
    if (date) {
      const day = date.getDate();
      const month = date.getMonth() + 1;
      const year = date.getFullYear();
      return `${day.toString().padStart(2, "0")}-${month
        .toString()
        .padStart(2, "0")}-${year}`;
    }
    return "";
  };

  return (
    <div className="booking-form-container">
      <h2>Book Now</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Select Option:
          <select
            value={selectedOption}
            onChange={(e) => setSelectedOption(e.target.value)}
          >
            <option value="">Select Branch</option>
            {dropdownOptions.map((option) => (
              <option key={option.ID} value={option.ID}>
                {option.Branch_name}
              </option>
            ))}
          </select>
        </label>
        <br />
        <label>
          Booking Date:
          <DatePicker
            placeholderText="Select-Date"
            selected={selectedDate}
            onChange={(date) => setSelectedDate(date)}
            dateFormat="dd-MM-yyyy"
          />
        </label>
        <br />
        <button type="submit" onClick={Postdata}>
          Submit
        </button>
        <Link to={"/"}>
          <button type="button">Close</button>
        </Link>
      </form>
      <div>
        <availableSlot/>
          
      </div>
    </div>
  );
}

export default BookingForm;
