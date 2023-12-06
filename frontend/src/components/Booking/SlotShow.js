import React, { useState, useEffect } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import { RiCheckboxBlankCircleFill } from "react-icons/ri";
import { GrPrevious } from "react-icons/gr";
import { GrNext } from "react-icons/gr";

import './SlotShow.css';

const SlotShow = (props) => {
  const [selectedDate, setSelectedDate] = useState(new Date());
  const [selectedSlots, setSelectedSlots] = useState([]);
  const [isDatePickerOpen, setIsDatePickerOpen] = useState(false);
  const [branchId, setBranchId] = useState(props.BranchId);
  const [availableSlots, setAvailableSlots] = useState([]);

  useEffect(() => {
    setBranchId(props.BranchId);
  }, [props.BranchId]);

  useEffect(() => {
    const fetchSlotsFromAPI = async () => {
      try {
        const payload = JSON.stringify({
          "branch_id": parseInt(branchId),
          "date": formatDate(selectedDate)
        });

        const response = await fetch("http://127.0.0.1:8080/user/get/avl/slots", {
          method: "POST",
          body: payload,
        });

        if (!response.ok) {
          throw new Error("Network response was not ok");
        }

        const responseJson = await response.json();

        if (responseJson && responseJson.available_slots) {
          setAvailableSlots(responseJson.available_slots);
        } else {
          console.error("Response is missing the expected 'available_slots' property:", responseJson);
        }
      } catch (error) {
        console.error("Error while fetching data:", error);
      }
    };

    fetchSlotsFromAPI();
  }, [selectedDate, branchId]);

  const handleDateChange = (date) => {
    setSelectedDate(date);
    setIsDatePickerOpen(false);
  };

  const handlePrevClick = () => {
    const newDate = new Date(selectedDate);
    newDate.setMonth(newDate.getMonth() - 1);
    setSelectedDate(newDate);
  };

  const handleNextClick = () => {
    const newDate = new Date(selectedDate);
    newDate.setMonth(newDate.getMonth() + 1);
    setSelectedDate(newDate);
  };

  const handleSlotSelect = (slot) => {
    console.log('Selected Slot:', slot);
  };

  const formatDate = (date) => {
    if (date) {
      const day = date.getDate();
      const month = date.getMonth() + 1;
      const year = date.getFullYear();
      return `${day.toString().padStart(2, "0")}-${month.toString().padStart(2, "0")}-${year}`;
    }
    return "";
  };

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
      <div className='row d-flex'>
        <div className='col-sm-12 col-md-3  border rounded-3 shadow-lg mt-2 text-center me-2'>
          <div className='mt-2 '>
            <GrPrevious className='fs-5 text-dark me-3' onClick={handlePrevClick} />
            <span className="me-3" onClick={() => setIsDatePickerOpen(!isDatePickerOpen)}>
              {selectedDate.toDateString()}
            </span>
            <GrNext className='fs-5 text-dark me-3' onClick={handleNextClick} />
          </div>
        </div>

        <div className="col-md-5 col-sm-12 mt-2">
          <div className='row border rounded-3 shadow-lg'>
            <div className="col">
              <p className='bookSlots'>
                <RiCheckboxBlankCircleFill className='fs-6 me-2 text-danger' /> Booked
              </p>
            </div>
            <div className="col">
              <p className='bookSlots'>
                <RiCheckboxBlankCircleFill className='fs-6 me-2 text-light border border-dark rounded-circle' /> Available
              </p>
            </div>
            <div className="col">
              <p className='bookSlots'>
                <RiCheckboxBlankCircleFill className='fs-6 me-2 text-info' /> Waiting
              </p>
            </div>
          </div>
        </div>
      </div>

      {isDatePickerOpen && (
        <div className='date-picker-slot col-9'>
          <DatePicker
            selected={selectedDate}
            onChange={handleDateChange}
            withPortal
            inline
          />
        </div>
      )}
      <div className='row'>
        <div class="card mt-2 shadow-lg col-md-8 col-sm-12 mb-2" >


          <div class="card-body">
            {Object.keys(slotsByPackage).length > 0 ? (
              Object.keys(slotsByPackage).map((packageName) => (
                <div key={packageName}>
                  <div className=''>
                    <h2>{packageName}</h2></div>
                  <div className=''>
                    <div className='d-flex'>
                      {slotsByPackage[packageName].map((slot) => (
                        <div className='m-3'>

                          <div class="card rounded-3 shadow-lg" style={{ width: '150px', height: '100px' }} key={slot.id} onClick={() => handleSlotSelect(slot)}>
                            <div class="card-body">
                              <p class="card-title fs-6">{`${slot.Slot.Start_time} to ${slot.Slot.End_time}`}</p>
                              <h6 class="card-subtitle mb-2 text-muted">Price: {slot.Price}</h6>

                            </div>
                          </div>

                        </div>
                      ))}
                    </div>
                  </div>

                </div>
              ))
            ) : (
              <p>No available slots found.</p>
            )}
          </div>

        </div>
        <div className='col-md-4 col-12' ><div class="card border-success mb-3" >
          <div class="card-header bg-transparent border-success">Book Slot Price</div>
          <div class="card-body text-success">
            <h5 class="card-title">Login Please</h5>
            <p class="card-text">DBBDJKSJDFKSJK </p>
          </div>
          <div class="card-footer bg-transparent border-success">Footer</div>
        </div></div>
      </div>
    </div>


  );
};

export default SlotShow








// import React, { useState, useEffect } from 'react';
// import Calendar from 'react-calendar';
// import 'react-calendar/dist/Calendar.css';
// import './SlotShow.css';

// const SlotShow = (props) => {
//     const [selectedDate, setSelectedDate] = useState(new Date());
//     const [selectedSlots, setSelectedSlots] = useState([]);
//     const [isCalendarOpen, setIsCalendarOpen] = useState(false);
//     const [branchId, setbranchId] = useState([]);

//     useEffect(() => {
//         setbranchId(props.BranchId);
//     }, [props.BranchId]);
//     console.log(branchId);

//     useEffect(() => {
//         // Fetch your slots from the API here and update the selectedSlots state
//         // For example, you might use the fetch API or axios for this purpose
//         const fetchSlotsFromAPI = async () => {
//             try {
//                 const response = await fetch('your_api_endpoint');
//                 const data = await response.json();
//                 // Assuming your API returns an array of slots
//                 setSelectedSlots(data.slots);
//             } catch (error) {
//                 console.error('Error fetching slots:', error);
//             }
//         };

//         fetchSlotsFromAPI();
//     }, [selectedDate]); // Fetch slots whenever the selectedDate changes

//     const handleDateChange = (date) => {
//         setSelectedDate(date);
//         setIsCalendarOpen(false);
//         // Add your logic here for handling the selected date
//     };

//     const handlePrevClick = () => {
//         const newDate = new Date(selectedDate);
//         newDate.setMonth(newDate.getMonth() - 1);
//         setSelectedDate(newDate);
//     };

//     const handleNextClick = () => {
//         const newDate = new Date(selectedDate);
//         newDate.setMonth(newDate.getMonth() + 1);
//         setSelectedDate(newDate);
//     };

//     const handleSlotSelect = (slot) => {
//         // Handle slot selection logic
//         console.log('Selected Slot:', slot);
//     };

//     return (
//         <div>
//             <div>
//                 <button className='cal-btn' onClick={handlePrevClick}>Prev</button>
//                 <span onClick={() => setIsCalendarOpen(!isCalendarOpen)}>
//                     {selectedDate.toDateString()}
//                 </span>
//                 <button className='cal-btn' onClick={handleNextClick}>next</button>
//             </div>
//             {isCalendarOpen && (
//                 <div className='calender-slot'>
//                     <Calendar onChange={handleDateChange} value={selectedDate} /></div>
//             )}
//             <div>
//                 <label>Select Slots:</label>
//                 <div>
//                     {selectedSlots.map((slot) => (
//                         <button
//                             key={slot.id} // Replace with a unique identifier from your API
//                             onClick={() => handleSlotSelect(slot)}
//                             className={slot.isSelected ? 'selected' : ''}
//                         >
//                             {slot.time}
//                         </button>
//                     ))}
//                 </div>
//             </div>
//         </div>
//     );
// };

// export default SlotShow;



