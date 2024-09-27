import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import NavDropdown from "react-bootstrap/NavDropdown";

const navCss = {
  // background: "linear-gradient(to top right, #ffff99 50%, #66ffff 100%)"
};

function Header() {
  return (
    <Navbar expand="lg" className="bg-body-tertiary" style={navCss}>
      <Container>
        <Navbar.Brand href="#home">Tulsi Vault</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link href="#credential">Credentials</Nav.Link>
            <Nav.Link href="#addCredential">
              Add Credential / Media file
            </Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

export default Header;
