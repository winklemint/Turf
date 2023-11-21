import React, { useState, useEffect } from "react";

function SlotBooking({ availableSlots, loading, selectedBranch, selectedDate }) {
  const [checkedItems, setCheckedItems] = useState({});
  const [totalPrice, setTotalPrice] = useState(0);

  useEffect(() => {
    // Reset state when selectedBranch or selectedDate changes
    setCheckedItems({});
    setTotalPrice(0);
  }, [selectedBranch, selectedDate]);

  const handleCheckboxChange = (slotID, price) => {
    setCheckedItems((prevCheckedItems) => {
      const updatedCheckedItems = {
        ...prevCheckedItems,
        [slotID]: !prevCheckedItems[slotID],
      };

      // Calculate total price based on selected slots
      const newTotalPrice = Object.keys(updatedCheckedItems).reduce(
        (acc, key) => acc + (updatedCheckedItems[key] ? price : 0),
        0
      );

      setTotalPrice(newTotalPrice);
      return updatedCheckedItems;
    });
  };

  const handleBookNow = () => {
    // Implement the logic to navigate to the payment gateway with the total price and selected slots
    // Aap router ya koi bhi navigation method ka istemal kar sakte hain
    console.log("Booking Now! Total Price:", totalPrice);
    console.log("Selected Slots:", Object.keys(checkedItems).filter((slotID) => checkedItems[slotID]));
    // Payment gateway par jane ka logic implement karein
  };

  // Organize slots by package
  const slotsByPackage = availableSlots.reduce((acc, slot) => {
    const packageName = slot.Package;

    if (!acc[packageName]) {
      acc[packageName] = [];
    }

    acc[packageName].push(slot);

    return acc;
  }, {});

  return (
    <div>
      <h1>Booking Data</h1>
      {loading ? (
        <p>Loading...</p>
      ) : Object.keys(slotsByPackage).length > 0 ? (
        Object.keys(slotsByPackage).map((packageName) => (
          <div key={packageName}>
            <h2>{packageName}</h2>
            <table>
              <thead>
                <tr>
                  <th>Start Time</th>
                  <th>End Time</th>
                  <th>Price</th>
                  <th>Status</th>
                  <th>Book Slot</th>
                </tr>
              </thead>
              <tbody>
                {slotsByPackage[packageName].map((slot) => (
                  <tr key={slot.Slot.ID}>
                    <td>{slot.Slot.Start_time}</td>
                    <td>{slot.Slot.End_time}</td>
                    <td>{slot.Price}</td>
                    <td>
                      {slot.Is_booked === 2 ? (
                        <span style={{ color: "red" }}>Booked</span>
                      ) : null}
                    </td>
                    <td>
                      <input
                        type="checkbox"
                        checked={checkedItems[slot.Slot.ID] || false}
                        onChange={() => handleCheckboxChange(slot.Slot.ID, slot.Price)}
                      />
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        ))
      ) : (
        <p>No available slots found.</p>
      )}

      {Object.keys(checkedItems).length > 0 && (
        <div>
          <h3>Total Price: {totalPrice}</h3>
          <button onClick={handleBookNow}>Book Now</button>
        </div>
      )}
    </div>
  );
}

export default SlotBooking;










