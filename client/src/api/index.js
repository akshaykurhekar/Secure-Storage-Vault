import axios from "axios";

// get credential by vid
export const getCredentialByVaultId = async (id) =>{
    const result = await axios.get(`http://localhost:5000/get/credentials/${id}`);
    return result.data;
}

// get all vault
export const getAllVault = async () => {
  const aa = await axios.get("http://localhost:5000/get/vaults");
   return aa.data;
  };

// create vault
// {
//     "Name":"Personal Doc",
//     "Password":"999",
//     "Description":"doc...."
//   }
export const createVaultApi = (data) => {
    axios.post("http://localhost:5000/create/vault", data)
      .then((res) => {
        console.log(res.data);
        return res.data
      })
      .catch((err) => {
        console.log("ERROR in creating Vault:", err);
      });
  };

// create credential/
// {
//     "VId":1,
//     "Credential": {"Name":"GYM ID Card", "Cid":"08765436546"}
//   }
export const createCredential = async (data) => {
    const result = await axios.post("http://localhost:5000/create/credentials", data);
    return result.data;
  };

//update vault
// {
//     "Name":"Personal Doc 06",
//     "Password":"999",
//     "Desc":"doc...."
//   }
export const updateVault = async (id, data) => {
    axios.put(`http://localhost:5000/update/vault/${id}`, data)
      .then((res) => {
        // console.log(res.data);
        return res.data
      })
      .catch((err) => {
        console.log("ERROR in updating Vault:", err);
      });
  };

// update credential by id
// {
//     "VId":1,
//     "Credential": {"Name":"GYM ID Card", "Cid":"08765436546xx"}
//   }
export const updateCredential = async (id, data) => {
    axios.put(`http://localhost:5000/update/credential/${id}`, data)
      .then((res) => {
        // console.log(res.data);
        return res.data
      })
      .catch((err) => {
        console.log("ERROR in updating credential:", err);
      });
  };

// delete vault by id
export const deleteVault = async (id) => {
    axios.delete(`http://localhost:5000/delete/vault/${id}`)
      .then((res) => {
        // console.log(res.data);
        return res.data
      })
      .catch((err) => {
        console.log("ERROR in deleting Vault:", err);
      });
  };

// delete credential
export const deleteCredential = async (id) => {
    axios.delete(`http://localhost:5000/delete/credential/${id}`)
      .then((res) => {
        // console.log(res.data);
        return res.data
      })
      .catch((err) => {
        console.log("ERROR in deleting credential:", err);
      });
  };