import React,{useEffect,useState} from 'react';
import './App.css';
import Footer from './components/Footer';
import Section2 from './components/Section2';
import Header from './components/header';
import Section1 from './components/section1';
import BookNowButton from './components/BookNow';



function App() {
  const [Heading, setHeading] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8080/admin/heading/active')
          .then((response) => response.json())
          .then((data) => setHeading(data.data))
          .catch((error) => console.error('Error fetching Heading data:', error));
      }, []);
  return(
    <div className='body'>
      <Header/>
      <Section1 Headingdata ={Heading}></Section1>
      <Section2 Headingdata ={Heading}></Section2>
      <Footer Headingdata ={Heading}></Footer>
     </div>

    )
  
}

export default App;

