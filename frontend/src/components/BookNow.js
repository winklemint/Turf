import React, { useState } from 'react';
import { FaLocationDot } from 'react-icons/fa6';
import { FaStar } from 'react-icons/fa';
import { Swiper, SwiperSlide } from 'swiper/react';
import Footer from './Home/Footer';
import 'swiper/css';
import 'swiper/css/free-mode';
import 'swiper/css/navigation';
import 'swiper/css/thumbs';
import { FreeMode, Mousewheel, Navigation, Pagination, Thumbs } from 'swiper/modules';
import './BookNow.css'

import 'swiper/css';
import 'swiper/css/navigation';
import 'swiper/css/pagination';

function BookingForm() {
  const h2Style = {
    color: 'rgb(119 196 40)',
  };
  const [thumbsSwiper, setThumbsSwiper] = useState(null);
  return (
    <>
      <div className="container">
        <div className="row float-center">
          <div className="col-sm-12 col-md-4 col-lg-4 col-12">

            <Swiper
              style={{
                '--swiper-navigation-color': '#fff',
                '--swiper-pagination-color': '#fff',
                width: '100%', // Set the width as needed
                height: '400px', // Set the height as needed
              }}

              loop={true}
              spaceBetween={10}
              pagination={true}
              navigation={true}
              keyboard={true}
              thumbs={{ swiper: thumbsSwiper }}
              modules={[FreeMode, Navigation, Thumbs, Pagination, Mousewheel]}

              className="mySwiper2"
            >
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-1.jpg"
                  style={{ width: '100%', height: '100%' }}

                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-2.jpg"
                  style={{ width: '100%', height: '100%' }}

                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-3.jpg"
                  style={{ width: '100%', height: '100%' }}

                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-4.jpg"
                  style={{ width: '100%', height: '100%' }}

                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-5.jpg"
                  style={{ width: '100%', height: '100%' }}

                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-6.jpg"
                  style={{ width: '100%', height: '100%' }}

                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-7.jpg"
                  style={{ width: '100%', height: '100%' }}

                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-8.jpg"
                  style={{ width: '100%', height: '100%' }}

                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-9.jpg"
                  style={{ width: '100%', height: '100%' }}

                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-10.jpg"
                  style={{ width: '100%', height: '100%' }}

                />
              </SwiperSlide>
            </Swiper>
            {/* <img src="/assets/football-ground-flooring.jpg" className="w-100" /> */}
          </div>
          <div className="col-sm-12 col-md-8 col-lg-8 col-12 mt-5 text-center ">
            <div className="fs-6 p-2 float-start">
              <div className="container" style={{ marginTop: '50px' }}>
                <div className="row">
                  <h2 className=" fw-1 float-left branch-name" style={h2Style}>Walking Dreamz Turf | Rajendra Nagar
                  </h2>
                  <div className="col-1 location-sign">
                    <FaLocationDot className="text-danger fs-3 me-3" />
                  </div>
                  <div className="col-11">
                    <p className='branch-add'>
                      Rajendra Nagar, Near Rajendra Nagar Police station
                      Bijalpur, Indore (M.P.)
                    </p>
                  </div>
                  <div className="row text-start">
                    <div className="d-flex btn bg-warning text-light text-center col-2 badge text-wrap">
                      <FaStar className="me-2" /> 5
                    </div>
                    <div className="col-10 mb-1">
                      <div className="badge text-warning">
                        7 reviews / Write a review
                      </div>
                    </div>
                  </div>
                </div>
                <br />
                <div className="fs-6 text-start">Amenities:</div>
                <div className="fs-6 text-start">
                  <p
                    className="btn btn-secondary text-light me-2"
                    style={{ fontSize: '13px' }}
                  >
                    Seating
                  </p>
                  <p
                    className="btn btn-secondary text-light me-2"
                    style={{ fontSize: '13px' }}
                  >
                    Toilets
                  </p>
                  <p
                    className="btn btn-secondary text-light me-2"
                    style={{ fontSize: '13px' }}
                  >
                    Parking
                  </p>
                </div>
                <div className="text-start">
                  <button className="btn btn-outline-warning">Book Now</button>
                </div>
              </div>
            </div>
          </div>

        </div>
        <ul class="nav nav-tabs" id="myTab" role="tablist" style={{marginTop:'70px',}}>
          <li class="nav-item " role="presentation"style={{marginRight:'55px'}}>
            <button class="nav-link active" id="home-tab" data-bs-toggle="tab" data-bs-target="#home" type="button" role="tab" aria-controls="home" aria-selected="true">BOOKING</button>
          </li>
          <li class="nav-item" role="presentation"style={{marginRight:'55px'}}>
            <button class="nav-link" id="profile-tab" data-bs-toggle="tab" data-bs-target="#profile" type="button" role="tab" aria-controls="profile" aria-selected="false">MEMBERSHIP PLANE</button>
          </li>
          <li class="nav-item" role="presentation"style={{marginRight:'55px'}}>
            <button class="nav-link" id="contact-tab" data-bs-toggle="tab" data-bs-target="#contact" type="button" role="tab" aria-controls="contact" aria-selected="false">DETAILS</button>
          </li>
        </ul>
        <div class="tab-content" id="myTabContent">
          <div class="tab-pane fade show active" id="home" role="tabpanel" aria-labelledby="home-tab">HII <br/>I<br/>AM <br/>VISHAL</div>
          <div class="tab-pane fade" id="profile" role="tabpanel" aria-labelledby="profile-tab">...</div>
          <div class="tab-pane fade" id="contact" role="tabpanel" aria-labelledby="contact-tab">...</div>
        </div>


      </div>
      <Footer />


    </>
  );
}

export default BookingForm;



// import React, { useState, useEffect } from "react";
// import { Link } from "react-router-dom";
// import DatePicker from "react-datepicker";
// import "react-datepicker/dist/react-datepicker.css";
// import "./BookNow.css";
// import SlotBooking from "./Booking/SlotBooking";

// function BookingForm() {
//   const [selectedDate, setSelectedDate] = useState(null);
//   const [dropdownOptions, setDropdownOptions] = useState([]);
//   const [selectedOption, setSelectedOption] = useState("");
//   const [availableSlots, setAvailableSlots] = useState([]);
//   const [loading, setLoading] = useState(true);
//   const [showSlotBooking, setShowSlotBooking] = useState(false);
//   const [errors, setErrors] = useState({});

//   const Postdata = () => {
//     var payload = JSON.stringify({
//       "branch_id": parseInt(selectedOption),
//       "date": formatDate(selectedDate)
//     });

//     fetch("http://127.0.0.1:8080/user/get/avl/slots", {
//       method: "POST",
//       body: payload,
//     })

//       .then((res) => {
//         console.log(res);
//         if (!res.ok) {
//           throw new Error("Network response was not ok");
//         }
//         return res.json();
//       })
//       .then((responseJson) => {
//         // Check if the response has a "data" property
//         if (responseJson && responseJson.available_slots) {
//           console.log("success");
//           setAvailableSlots(responseJson.available_slots);
//           setLoading(false);
//           setShowSlotBooking(true);
//         } else {
//           console.error("Response is missing the expected 'data' property:", responseJson);
//         }
//       })
//       .catch((error) => {
//         console.error("Error while fetching data:", error);
//       });
//   };

//   useEffect(() => {
//     fetch("http://localhost:8080/admin/active/branch")
//       .then((res) => res.json())
//       .then((data) => {
//         setDropdownOptions(data.data);
//       });
//   }, []);

//   const handleChange = (e) => {
//     const { name, value } = e.target;
//     setSelectedOption(value);

//     // Validate selected option
//     if (!value) {
//       setErrors({ ...errors, selectedOption: "Please select a branch" });
//     } else {
//       setErrors({ ...errors, selectedOption: "" });
//     }
//   };

//   const handleSubmit = (e) => {
//     e.preventDefault();
//     console.log(JSON.stringify({
//       "branch_id": parseInt(selectedOption),
//       "date": formatDate(selectedDate)}));

//     // Validate form before submitting
//     if (!selectedOption || !selectedDate) {
//       setErrors({
//         selectedOption: !selectedOption ? "Please select a branch" : "",
//         selectedDate: !selectedDate ? "Please select a date" : "",
//       });
//       return;
//     }

//     // Call the Postdata function to fetch data and update availableSlots
//     Postdata();
//     setSelectedDate(null);
//     setSelectedOption("");
//   };

//   const formatDate = (date) => {
//     if (date) {
//       const day = date.getDate();
//       const month = date.getMonth() + 1;
//       const year = date.getFullYear();
//       return `${day.toString().padStart(2, "0")}-${month.toString().padStart(2, "0")}-${year}`;
//     }
//     return "";
//   };

//   return (
//     <div className="booking-form-container">
//       <h2>Book Now</h2>
//       <form onSubmit={handleSubmit}>
//         <label>
//           Select Option:
//           <select
//             name="selectedOption"
//             value={selectedOption}
//             onChange={handleChange}
//           >
//             <option value="">Select Branch</option>
//             {dropdownOptions.map((option) => (
//               <option key={option.ID} value={option.ID}>
//                 {option.Branch_name}
//               </option>
//             ))}
//           </select>
//         </label>
//         {errors.selectedOption && (
//           <div className="error-message">{errors.selectedOption}</div>
//         )}
//         <br />
//         <label>
//           Booking Date:
//           <DatePicker
//             placeholderText="Select Date"
//             selected={selectedDate}
//             onChange={(date) => setSelectedDate(date)}
//             dateFormat="dd-MM-yyyy"
//           />
//         </label>
//         {errors.selectedDate && (
//           <div className="error-message">{errors.selectedDate}</div>
//         )}
//         <br />
//         <button type="submit">Submit</button>
//         <Link to={"/"}>
//           <button type="button">Close</button>
//         </Link>
//       </form>
//       {showSlotBooking && (
//         <SlotBooking availableSlots={availableSlots} loading={loading} />
//       )}
//     </div>
//   );
// }

// export default BookingForm;


