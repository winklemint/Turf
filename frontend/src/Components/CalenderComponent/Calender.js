import React, { useState } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';

const Calendar = () => {
  const [selectedDate, setSelectedDate] = useState(null);

  // Function to disable specific dates
  const isWeekday = (date) => {
    const day = date.getDay();
    return day !== 0 && day !== 6; // Disable weekends (Sunday = 0, Saturday = 6)
  };

  return (
    <div>
      <h2>Calendar</h2>
      <DatePicker
        selected={selectedDate}
        onChange={date => setSelectedDate(date)}
        locale="en"
        dateFormat="yyyy-MM-dd"
        filterDate={isWeekday}
      />
    </div>
  );
};

export default Calendar;
