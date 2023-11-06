import React, { useState, useEffect } from "react";

function SlotBooking() {
  
 
  return (
    <div>
      <h1>Booking Data</h1>
      {loading ? (
        <p>Loading...</p>
      ) : (
        <ul>
          {availableSlots.map((slot) => (
            <li key={slot.id}>
              <span>Start Time: {slot.start_time}</span>
              <span>End Time: {slot.end_time}</span>
              {slot.Is_booked === 1 ? (
                <input type="checkbox" checked />
              ) : slot.Is_booked === 2 ? (
                <span style={{ color: "red" }}>Booked</span>
              ) : null}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default SlotBooking;
