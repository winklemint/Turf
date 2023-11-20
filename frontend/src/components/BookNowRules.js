import React from 'react';
import { FaStar } from 'react-icons/fa';

function BookNowRules() {
  const h2Style = {
    color: 'rgb(119 196 40)',
  };

  const stars = Array(5).fill(null);

  return (
    <div className="container">
      <h2 className="mt-3 me-4" style={h2Style}>
        Rules
      </h2>
      <div>
        <b>
          - NO SMOKING <br />
          - NO ALCOHOL CONSUMPTION <br />
          - Book Ground 1 & Ground 2 for full ground <br />
          - NO MORE THAN 20 PEOPLE PER BOOKING <br />
          - NO OUTSIDE FOOD & EATABLES ALLOWED <br />- FOR EVENT BOOKINGS PLEASE
          DROP A WHATSAPP MESSAGE BY CLICKING: 8169729906
        </b>
        <br />
      </div>
      <div className="mt-4">
        <h2 style={h2Style}>Reviews</h2>
        <div className="mt-2 ">
          <div className="card p-2 border border-rounded mb-2">
            <p>
              <b>Anan Sharma</b>
              <br />
              <i class="bi bi-quote text-secondary">
                Ordered Friday morning, delivered Saturday morning in the exact
                spot on the drive we asked for. Very polite and helpful driver
                dealing patiently with my 82 y.o dad. He didn’t catch his name
                but hopefully you can trace him and pass on my thanks.
              </i>
              <br />
              {stars.map((_, index) => (
                <FaStar key={index} className="fs-5 text-warning" />
              ))}
            </p>
          </div>
          <div className="card p-2 border border-rounded  mb-2">
            <p>
              <b>Abhishek Ved</b>
              <br />
              <i class="bi bi-quote text-secondary">
                Ordered Friday morning, delivered Saturday morning in the exact
                spot on the drive we asked for. Very polite and helpful driver
                dealing patiently with my 82 y.o dad. He didn’t catch his name
                but hopefully you can trace him and pass on my thanks.
              </i>
              <br />
              {stars.map((_, index) => (
                <FaStar key={index} className="fs-5 text-warning" />
              ))}
            </p>
          </div>
          <div className="card p-2 border border-rounded  mb-2">
            <p>
              <b>Nishant Gupta</b>
              <br />
              <i class="bi bi-quote text-secondary">
                Ordered Friday morning, delivered Saturday morning in the exact
                spot on the drive we asked for. Very polite and helpful driver
                dealing patiently with my 82 y.o dad. He didn’t catch his name
                but hopefully you can trace him and pass on my thanks.
              </i>
              <br />
              {stars.map((_, index) => (
                <FaStar key={index} className="fs-5 text-warning" />
              ))}
            </p>
          </div>
          <div className="card p-2 border border-rounded  mb-2">
            <p>
              <b>Rohit Arya</b>
              <br />
              <i class="bi bi-quote text-secondary">
                Ordered Friday morning, delivered Saturday morning in the exact
                spot on the drive we asked for. Very polite and helpful driver
                dealing patiently with my 82 y.o dad. He didn’t catch his name
                but hopefully you can trace him and pass on my thanks.
              </i>
              <br />
              {stars.map((_, index) => (
                <FaStar key={index} className="fs-5 text-warning" />
              ))}
            </p>
          </div>
          <div className="card p-2 border border-rounded  mb-2">
            <p>
              <b>Geeta Morna</b>
              <br />
              <i class="bi bi-quote text-secondary">
                Ordered Friday morning, delivered Saturday morning in the exact
                spot on the drive we asked for. Very polite and helpful driver
                dealing patiently with my 82 y.o dad. He didn’t catch his name
                but hopefully you can trace him and pass on my thanks.
              </i>
              <br />
              {stars.map((_, index) => (
                <FaStar key={index} className="fs-5 text-warning" />
              ))}
            </p>
          </div>
          <div className="card p-2 border border-rounded  mb-2">
            <p>
              <b>Raj Sharma</b>
              <br />
              <i class="bi bi-quote text-secondary">
                Ordered Friday morning, delivered Saturday morning in the exact
                spot on the drive we asked for. Very polite and helpful driver
                dealing patiently with my 82 y.o dad. He didn’t catch his name
                but hopefully you can trace him and pass on my thanks.
              </i>
              <br />
              {stars.map((_, index) => (
                <FaStar key={index} className="fs-5 text-warning" />
              ))}
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default BookNowRules;
