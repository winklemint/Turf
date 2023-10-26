import React, { useState } from 'react';
import Calendar from '../CalenderComponent/Calender';

const Booking = () => {
  const [value, setValue] = useState('');
  
  // Function to disable specific dates
  const dateDisabled = (ymd, date) => {
    // Disable weekends (Sunday = 0, Saturday = 6) and days that fall on the 13th of the month
    const weekday = date.getDay();
    const day = date.getDate();
    // Return true if the date should be disabled
    return weekday === 0 || weekday === 6 || day === 13;
  };

  // Calculate the min and max dates
  const now = new Date();
  const today = new Date(now.getFullYear(), now.getMonth(), now.getDate());
  const minDate = new Date(today);
  minDate.setMonth(minDate.getMonth() - 2);
  minDate.setDate(15);
  const maxDate = new Date(today);
  maxDate.setMonth(maxDate.getMonth() + 2);
  maxDate.setDate(15);

  return (
    <div>
      <div>Booking</div>
      {/* Calendar with dateDisabled function */}
      <div>
        <Calendar
          value={value}
          dateDisabledFn={dateDisabled}
          locale="en"
        />
      </div>
    </div>
  );
};

export default Booking;
