import "./App.css";
import { Button, Col, Container, Row } from "react-bootstrap";
import Header from "./component/navbar";
import BoxCard from "./component/card";
import { useEffect, useState } from "react";
import axios from "axios";
import AddCredential from "./component/addCredentialModal";
import AddVault from "./component/addVaultModal";
import { getAllVault } from "./api";

function App() {
  const [users, setUsers] = useState([]);
  const [showModel, setShowModel] = useState(false);

  useEffect( () => {
    axios.get("http://localhost:5000/get/vaults").then((result)=>{
      console.log(result.data);
      setUsers(result.data);
    })
    console.log("VAULT::",users)
    // setUsers(aa)
  }, []);

  const getCredential = async () => {
    axios.get("http://localhost:5000/get/vaults").then((result)=>{
      console.log(result.data);
      setUsers(result.data);
    })
    console.log("VAULTS LIST:",users);
  };


  return (
    <div className="App">
      <Header />
      <div style={{ padding: "10px" }}>
        <Button style={{ margin: "10px" }} onClick={getCredential}>
          Get All Vaults
        </Button>
        <AddCredential />
        <AddVault />
      </div>
      <Container>
        <Col style={{ marginTop: "1rem", padding: "1rem" }}>
          { users !== null ? (
            users.map((item) => {
              return (
                <Row key={item.id} style={{ marginTop: "1rem" }}>
                  {" "}
                  <BoxCard
                    key={item.id}
                    name={item.name}
                    email={item.email}
                    id={item.id}
                    password={item.password}
                  />
                </Row>
              );
            })
          ) : (
            <p>No credential Found</p>
          )}
        </Col>
      </Container>
    </div>
  );
}

export default App;
