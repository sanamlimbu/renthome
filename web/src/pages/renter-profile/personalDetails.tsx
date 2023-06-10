import {
  Box,
  Button,
  FormControl,
  MenuItem,
  Select,
  TextField,
  Typography,
  useMediaQuery,
} from "@mui/material";
import { useContext, useEffect } from "react";
import { Controller, SubmitHandler, useForm } from "react-hook-form";
import { API_ADDRESS } from "../../config";
import { UserContext } from "../../context/user";
import { getTokenFromLocalStorage } from "../../helpers/auth";

interface IFormInput {
  firstName: string;
  lastName: string;
  phoneNumber: string;
  day: string;
  month: string;
  year: string;
}

export default function PersonalDetails() {
  const matchesSm = useMediaQuery("(max-width: 959px)");
  const { user } = useContext(UserContext);

  useEffect(() => {
    (async function () {
      try {
        const res = await fetch(`${API_ADDRESS}/users/?userID:${user?.id}`, {
          method: "POST",
          headers: {
            Authorization: `Bearer ${getTokenFromLocalStorage()}`,
          },
        });

        if (res.ok) {
          const data = await res.json();
        }
      } catch (error) {}
    })();
  }, []);

  const { control, handleSubmit, register, reset } = useForm({
    defaultValues: {
      firstName: "",
      lastName: "",
      phoneNumber: "",
      day: "Date",
      month: "Month",
      year: "Year",
    },
  });

  // Set days and months range
  const days = Array.from({ length: 31 }, (_, i) => i + 1);
  const months = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];

  // Set years range
  const currentYear = new Date().getFullYear();
  const startYear = currentYear - 112;
  const years: number[] = [];
  for (let year = startYear; year <= currentYear; year++) {
    years.push(year);
  }

  const onSubmit: SubmitHandler<IFormInput> = (data) => {
    console.log(data);
  };

  return (
    <Box
      sx={{
        maxWidth: "34em",
      }}
    >
      <Typography fontWeight={"bold"}>RENTER PROFILE</Typography>
      <Typography variant="h1" sx={{ marginBottom: "1em" }}>
        Personal Details
      </Typography>
      <form
        style={{ display: "flex", flexDirection: "column", gap: "1em" }}
        onSubmit={handleSubmit(onSubmit)}
      >
        <div>
          <Typography fontWeight={"600"}>First name</Typography>
          <TextField
            variant="outlined"
            value={"Sanam"}
            fullWidth
            {...register("firstName")}
          />
        </div>
        <div>
          <Typography fontWeight={"600"}>Last name</Typography>
          <TextField
            variant="outlined"
            value={"Limbu"}
            fullWidth
            {...register("lastName")}
          />
        </div>
        <div>
          <Typography fontWeight={"600"}>Date of birth</Typography>
          <div
            style={{
              display: "flex",
              gap: "20px",
              justifyContent: "space-between",
            }}
          >
            <FormControl fullWidth {...register("day")}>
              <Controller
                name="day"
                control={control}
                render={({ field }) => (
                  <Select {...field}>
                    <MenuItem value={"Date"}>Date</MenuItem>
                    {days.map((day) => (
                      <MenuItem key={"day-" + day} value={day}>
                        {day}
                      </MenuItem>
                    ))}
                  </Select>
                )}
              />
            </FormControl>
            <FormControl fullWidth {...register("month")}>
              <Controller
                name="month"
                control={control}
                render={({ field }) => (
                  <Select {...field}>
                    <MenuItem value={"Month"}>Month</MenuItem>
                    {months.map((month) => (
                      <MenuItem key={month} value={month}>
                        {month}
                      </MenuItem>
                    ))}
                  </Select>
                )}
              />
            </FormControl>
            <FormControl fullWidth {...register("year")}>
              <Controller
                name="year"
                control={control}
                render={({ field }) => (
                  <Select {...field}>
                    <MenuItem value={"Year"}>Year</MenuItem>
                    {years.map((year) => (
                      <MenuItem key={year} value={year}>
                        {year}
                      </MenuItem>
                    ))}
                  </Select>
                )}
              />
            </FormControl>
          </div>
        </div>
        <div>
          <Typography fontWeight={"600"}>
            Phone number (mobile prefered)
          </Typography>
          <TextField
            variant="outlined"
            value={"0299910121"}
            fullWidth
            {...register("phoneNumber")}
          />
          <Typography fontSize={"15px"}>
            Use numbers only, without spaces or other characters, e.g.
            0416222333 or 0244443333.
          </Typography>
        </div>
        <div
          style={{
            display: "flex",
            flexDirection: matchesSm ? "column" : "row",
            gap: "1em",
          }}
        >
          <Button sx={{ fontWeight: "bold" }} fullWidth>
            Cancel
          </Button>
          <Button
            variant="contained"
            sx={{ fontWeight: "bold" }}
            type="submit"
            fullWidth
          >
            Save and back
          </Button>
        </div>
      </form>
    </Box>
  );
}
function useState<T>(): IFormInput {
  throw new Error("Function not implemented.");
}
