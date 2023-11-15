

// import React, { useState } from "react";

// function SlotBooking({ availableSlots, loading }) {
//   const [checkedItems, setCheckedItems] = useState({});

//   const handleCheckboxChange = (slotID) => {
//     setCheckedItems((prevCheckedItems) => ({
//       ...prevCheckedItems,
//       [slotID]: !prevCheckedItems[slotID],
//     }));
//   };

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
//               <input
//                 type="checkbox"
//                 checked={checkedItems[slot.Slot.ID] || false}
//                 onChange={() => handleCheckboxChange(slot.Slot.ID)}
//               />
//               {slot.Is_booked === 2 ? (
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
  const [selectedSlot, setSelectedSlot] = useState(null);

  const handleCheckboxChange = (slotID) => {
    setCheckedItems((prevCheckedItems) => ({
      ...prevCheckedItems,
      [slotID]: !prevCheckedItems[slotID],
    }));
  };

  const handleBookNow = (slot) => {
    setSelectedSlot(slot);
  };

  console.log("availableSlots:", availableSlots);
  console.log("availableSlots.data:", availableSlots.data);

  if (!availableSlots || !availableSlots.data) {
    return <p>No available slots found.</p>;
  }

  const platinumSlots = availableSlots.data.filter((slot) => slot.Package === "Platinum") || [];
  const goldSlots = availableSlots.data.filter((slot) => slot.Package === "Gold") || [];

  return (
    <div>
      <h1>Booking Data</h1>
      {loading ? (
        <p>Loading...</p>
      ) : (
        <>
          {platinumSlots.length > 0 && (
            <>
              <h2>Platinum Packages</h2>
              <ul>
                {platinumSlots.map((slot) => (
                  <li key={slot.Slot.ID}>
                    <span>Start Time: {slot.Slot.Start_time}</span>
                    {/* Other slot details */}
                    <input
                      type="checkbox"
                      checked={checkedItems[slot.Slot.ID] || false}
                      onChange={() => handleCheckboxChange(slot.Slot.ID)}
                    />
                    {slot.Is_booked === 2 ? (
                      <span style={{ color: "red" }}>Booked</span>
                    ) : (
                      <button onClick={() => handleBookNow(slot)}>
                        Book Now
                      </button>
                    )}
                  </li>
                ))}
              </ul>
            </>
          )}
          {goldSlots.length > 0 && (
            <>
              <h2>Gold Packages</h2>
              <ul>
                {goldSlots.map((slot) => (
                  <li key={slot.Slot.ID}>
                    <span>Start Time: {slot.Slot.Start_time}</span>
                    {/* Other slot details */}
                    <input
                      type="checkbox"
                      checked={checkedItems[slot.Slot.ID] || false}
                      onChange={() => handleCheckboxChange(slot.Slot.ID)}
                    />
                    {slot.Is_booked === 2 ? (
                      <span style={{ color: "red" }}>Booked</span>
                    ) : (
                      <button onClick={() => handleBookNow(slot)}>
                        Book Now
                      </button>
                    )}
                  </li>
                ))}
              </ul>
            </>
          )}

          {selectedSlot && (
            <div>
              <h2>Selected Slot Details</h2>
              <p>Start Time: {selectedSlot.Slot.Start_time}</p>
              {/* Display other details of the selected slot */}
              <button onClick={() => setSelectedSlot(null)}>
                Close Details
              </button>
            </div>
          )}
        </>
      )}
    </div>
  );
}

export default SlotBooking;


