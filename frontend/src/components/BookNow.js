import React, { useState } from "react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css"; // Import the date picker CSS

function BookingForm(props) {
  const [selectedDate, setSelectedDate] = useState(null); // State to store the selected date

  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle the form submission, e.g., send data to the server
    console.log("Name:", e.target.name.value);
    console.log("Mobile Number:", e.target.mobile.value);
    console.log("Booking Date (dd/mm/yyyy):", formatDate(selectedDate));
    // Close the form
    props.onClose();
  };

  const formatDate = (date) => {
    if (date) {
      const day = date.getDate();
      const month = date.getMonth() + 1; // Months are 0-indexed
      const year = date.getFullYear();
      return `${day.toString().padStart(2, "0")}-${month.toString().padStart(2, "0")}-${year}`;
    }
    return "";
  };

  return (
    <div>
      <h2>Book Now</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Name:
          <input type="text" name="name" />
        </label>
        <br />
        <label>
          Mobile Number:
          <input type="text" name="mobile" />
        </label>
        <br />
        <label>
          Booking Date:
          <DatePicker
            selected={selectedDate}
            onChange={(date) => setSelectedDate(date)}
            dateFormat="dd/MM/yyyy"
          />
        </label>
        <br />
        <button type="submit">Submit</button>
        <button onClick={props.onClose}>Close</button>
      </form>
    </div>
  );
}

export default BookingForm;
