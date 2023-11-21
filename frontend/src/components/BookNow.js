import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import "./BookNow.css";
import SlotBooking from "./Booking/SlotBooking";

function BookingForm() {
  const [selectedDate, setSelectedDate] = useState(null);
  const [dropdownOptions, setDropdownOptions] = useState([]);
  const [selectedOption, setSelectedOption] = useState("");
  const [availableSlots, setAvailableSlots] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showSlotBooking, setShowSlotBooking] = useState(false);
  const [errors, setErrors] = useState({});

  const Postdata = () => {
    var payload = JSON.stringify({
      "branch_id": parseInt(selectedOption),
      "date": formatDate(selectedDate)
    });

    fetch("http://127.0.0.1:8080/user/get/avl/slots", {
      method: "POST",
      body: payload,
    })

      .then((res) => {
        console.log(res);
        if (!res.ok) {
          throw new Error("Network response was not ok");
        }
        return res.json();
      })
      .then((responseJson) => {
        // Check if the response has a "data" property
        if (responseJson && responseJson.available_slots) {
          console.log("success");
          setAvailableSlots(responseJson.available_slots);
          setLoading(false);
          setShowSlotBooking(true);
        } else {
          console.error("Response is missing the expected 'data' property:", responseJson);
        }
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

  const handleChange = (e) => {
    const { name, value } = e.target;
    setSelectedOption(value);

    // Validate selected option
    if (!value) {
      setErrors({ ...errors, selectedOption: "Please select a branch" });
    } else {
      setErrors({ ...errors, selectedOption: "" });
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(JSON.stringify({
      "branch_id": parseInt(selectedOption),
      "date": formatDate(selectedDate)}));

    // Validate form before submitting
    if (!selectedOption || !selectedDate) {
      setErrors({
        selectedOption: !selectedOption ? "Please select a branch" : "",
        selectedDate: !selectedDate ? "Please select a date" : "",
      });
      return;
    }

    // Call the Postdata function to fetch data and update availableSlots
    Postdata();
    setSelectedDate(null);
    setSelectedOption("");
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

  return (
    <div className="booking-form-container">
      <h2>Book Now</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Select Option:
          <select
            name="selectedOption"
            value={selectedOption}
            onChange={handleChange}
          >
            <option value="">Select Branch</option>
            {dropdownOptions.map((option) => (
              <option key={option.ID} value={option.ID}>
                {option.Branch_name}
              </option>
            ))}
          </select>
        </label>
        {errors.selectedOption && (
          <div className="error-message">{errors.selectedOption}</div>
        )}
        <br />
        <label>
          Booking Date:
          <DatePicker
            placeholderText="Select Date"
            selected={selectedDate}
            onChange={(date) => setSelectedDate(date)}
            dateFormat="dd-MM-yyyy"
          />
        </label>
        {errors.selectedDate && (
          <div className="error-message">{errors.selectedDate}</div>
        )}
        <br />
        <button type="submit">Submit</button>
        <Link to={"/"}>
          <button type="button">Close</button>
        </Link>
      </form>
      {showSlotBooking && (
        <SlotBooking availableSlots={availableSlots} loading={loading} />
      )}
    </div>
  );
}

export default BookingForm;

// import React, { useRef, useState } from "react";
// // Import Swiper React components
// import { Swiper, SwiperSlide } from "swiper/react";

// // Import Swiper styles
// import "swiper/css";
// import "swiper/css/pagination";
// import "swiper/css/navigation";

// // import './styles.css';

// // import required modules
// import { Pagination, Navigation } from "swiper/modules";

// function BookingForm() {
//   return (
//     <div className="container">
//   <div className="row">
//     <div className="col-md-5">
//       <Swiper
//         slidesPerView={1}
//         spaceBetween={30}
//         loop={true}
//         pagination={{
//           clickable: true,
//         }}
//         navigation={true}
//         modules={[Pagination, Navigation]}
//         className="mySwiper"
//       >
//         <SwiperSlide>Slide 1</SwiperSlide>
//         <SwiperSlide>Slide 2</SwiperSlide>
//         {/* ... (up to Slide 9) */}
//       </Swiper>
//     </div>

//     <div className="col-md-7">
//       <div className="d-flex flex-column justify-content-between h-100 pb--50-mob">
//         <div className="row">
//           <div className="col-8">
//             <div className='d-flex justify-content-between align-items-center mt--5 mt--10-mob mb--15 mb--10-mob'>
//               <div className="style_exclusive-badge__338Sr">
//                 <div className="style_round-wrapper__1Pm1_">T</div>
//                 "EXCLUSIVE"
//               </div>
//             </div>
//           </div>
//         </div>
//       </div>
//     </div>
//   </div>
// </div>

//   );
// }
// export default BookingForm;
