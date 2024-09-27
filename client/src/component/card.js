import Button from "react-bootstrap/Button";
import Card from "react-bootstrap/Card";
import { Row, Col, Container } from "react-bootstrap";
import { useState } from "react";
import CredentialCard from "./credentialCard";
import {getCredentialByVaultId} from "../api"
import axios from "axios";

const cardStyle = {
  background: "linear-gradient(to bottom left, #000009 0%, #303350 58%)",
  // background:"#ff6b9f",
  width: "100%",
  justifyContent: "center",
  border: "solid green 1px",
  //  backgroundColor:"#e6fff7"
  color: "white",
};

function BoxCard({ name, email, password, id }) {
  // console.log(id)
  const [unlockButton, setUnlockButton] = useState(false);
  const [unlock, setUnlock] = useState(false);
  const [Credentials, setCredentials] = useState([]);

  const unlockVault = (pass) => {
    if (password === pass) {
      setUnlock(true);
      console.log("Pass match");
      showCredential(unlock, id);
      setUnlockButton(false);
    }
  };

  const showCredential = async (status, id) => {
    console.log("Show credential...:", status, id);
    const result = await getCredentialByVaultId(id);
    setCredentials(result);

    // axios
    //   .get(`http://localhost:4000/credential/${id}`)
    //   .then((res) => {
    //     console.log(res.data);
    //     setCredentials(res.data);
    //   })
    //   .catch((err) => {
    //     console.log("ERROR in getting Credential list:", err);
    //   });
    console.log(Credentials);
  };

  // this card is to show list vaults
  return (
    <Card style={cardStyle} key={id}>
      <Card.Body flex flexDirection="row">
        <Row>
          <Col style={{ color: "white" }}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="28"
              height="26"
              fill="currentColor"
              class="bi bi-wallet2"
              viewBox="0 0 16 16"
            >
              <path d="M12.136.326A1.5 1.5 0 0 1 14 1.78V3h.5A1.5 1.5 0 0 1 16 4.5v9a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 0 13.5v-9a1.5 1.5 0 0 1 1.432-1.499zM5.562 3H13V1.78a.5.5 0 0 0-.621-.484zM1.5 4a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5z" />
            </svg>
          </Col>{" "}
          |
          <Col>
            <Card.Title> {name}</Card.Title>
          </Col>
          |{" "}
          <Col>
            <Card.Text>{email}</Card.Text>
          </Col>
        </Row>
      </Card.Body>
      <hr></hr>
      <Container>
        <Row style={{ borderColor: "black", padding: "1rem" }}>
          <Button
            style={{ border: "solid pink 1px" }}
            variant="success"
            size="sm"
            onClick={() => {
              setUnlockButton(!unlockButton);
            }}
          >
            Unlock Vault
          </Button>

          {unlockButton ? (
            <input
              placeholder="enter pass"
              onChange={(e) => {
                unlockVault(e.target.value);
              }}
            />
          ) : null}
          {Credentials != null
            ? Credentials.map((item) => {
                return (
                  <Col key={item.id} style={{ marginTop: "1rem" }}>
                    {" "}
                    <CredentialCard
                      key={item.id}
                      name={item.credential.name}
                      uid={item.uid}
                      cid={item.credential.cid}
                    />{" "}
                  </Col>
                );
              })
            : null}
        </Row>
      </Container>
    </Card>
  );
}

export default BoxCard;
