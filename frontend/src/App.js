import './App.css';
import Footer from './components/Footer';
import Section2 from './components/Section2';
import Header from './components/header';
import BookNow from './components/BookNow';
import Section1 from './components/section1';

function App() {
  return (
    <div className="body">
      <Header />
      <Section1 />
      <Section2 />
      <Footer />
      <BookNow />
    </div>
  );
}

export default App;
