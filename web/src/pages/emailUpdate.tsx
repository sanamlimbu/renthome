import { EmailOutlined } from "@mui/icons-material";
import {
  Alert,
  Button,
  Card,
  InputAdornment,
  TextField,
  Typography,
} from "@mui/material";
import { useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import RentHomeLogo from "../components/rentHomeLogo";
import { API_ADDRESS } from "../config";
import { getTokenFromLocalStorage } from "../helpers/auth";
import "../styles/index.css";
import { ErrorResponse } from "../types/types";

interface IEmailUpdateInput {
  new_email: string;
}

export default function EmailUpdatePage() {
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<IEmailUpdateInput>();
  const navigate = useNavigate();
  const [error, setError] = useState("");

  const onSubmit: SubmitHandler<IEmailUpdateInput> = async (input) => {
    try {
      const res = await fetch(`${API_ADDRESS}/auth/email-update`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${getTokenFromLocalStorage()}`,
        },
        body: JSON.stringify(input),
      });
      if (res.ok) {
        const data = res.json();
        navigate("/me");
      } else {
        const data: ErrorResponse = await res.json();
        setError(data.message);
      }
    } catch (error) {}
  };

  return (
    <div className="Container-centered">
      <Card variant="outlined" className="Card">
        <RentHomeLogo />
        <Typography textAlign={"center"} fontWeight={"700"} fontSize={"18px"}>
          Update your email address
        </Typography>
        <form onSubmit={handleSubmit(onSubmit)} className="Form">
          {error && <Alert severity="warning">{error}</Alert>}
          <TextField
            variant="outlined"
            placeholder="New email address"
            {...register("new_email", {
              required: {
                value: true,
                message: "Please enter a valid email address.",
              },
            })}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <EmailOutlined />
                </InputAdornment>
              ),
            }}
          />
          {errors.new_email && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.new_email.message}
            </span>
          )}
          <Button
            variant="contained"
            size="large"
            sx={{ textTransform: "none", fontWeight: "600" }}
            type="submit"
          >
            Next
          </Button>
        </form>
        <div style={{ paddingTop: "14px", paddingBottom: "14px" }}>
          <Button
            sx={{ textTransform: "none", fontWeight: "700" }}
            fullWidth
            onClick={() => navigate("/me")}
          >
            Cancel
          </Button>
        </div>
      </Card>
    </div>
  );
}
