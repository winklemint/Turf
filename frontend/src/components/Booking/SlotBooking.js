
// import React from "react";

// function SlotBooking({ availableSlots, loading }) {
//   console.log("availableSlots:", availableSlots);
//   console.log("availableSlots.data:", availableSlots.data);
//   return (
//     <div>
//       <h1>Booking Data</h1>
//       {loading ? (
//         <p>Loading...</p>
//       ) : availableSlots && availableSlots.data ? (
//         <ul>
//           {availableSlots.data.map((slot) => (
//             <li key={slot.Slot.ID}>
//               <span>Start Time: {slot.Slot.Start_time}</span>
//               <span>End Time: {slot.Slot.End_time}</span>
//               {slot.Is_booked === 1 ? (
//                 <input type="checkbox" checked />
//               ) : slot.Is_booked === 2 ? (
//                 <span style={{ color: "red" }}>Booked</span>
//               ) : null}
//             </li>
//           ))}
//         </ul>
//       ) : (
//         <p>No available slots found.</p>
//       )}
//     </div>
//   );
  
// }

// export default SlotBooking;
import React, { useState } from "react";

function SlotBooking({ availableSlots, loading }) {
  const [checkedItems, setCheckedItems] = useState({});

  const handleCheckboxChange = (slotID) => {
    setCheckedItems((prevCheckedItems) => ({
      ...prevCheckedItems,
      [slotID]: !prevCheckedItems[slotID],
    }));
  };

  console.log("availableSlots:", availableSlots);
  console.log("availableSlots.data:", availableSlots.data);

  return (
    <div>
      <h1>Booking Data</h1>
      {loading ? (
        <p>Loading...</p>
      ) : availableSlots && availableSlots.data ? (
        <ul>
          {availableSlots.data.map((slot) => (
            <li key={slot.Slot.ID}>
              <span>Start Time: {slot.Slot.Start_time}</span>
              <span>End Time: {slot.Slot.End_time}</span>
              <input
                type="checkbox"
                checked={checkedItems[slot.Slot.ID] || false}
                onChange={() => handleCheckboxChange(slot.Slot.ID)}
              />
              {slot.Is_booked === 2 ? (
                <span style={{ color: "red" }}>Booked</span>
              ) : null}
            </li>
          ))}
        </ul>
      ) : (
        <p>No available slots found.</p>
      )}
    </div>
  );
}

export default SlotBooking;
