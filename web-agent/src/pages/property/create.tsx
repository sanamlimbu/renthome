import {
  Box,
  Button,
  FormControl,
  FormControlLabel,
  InputAdornment,
  MenuItem,
  Radio,
  RadioGroup,
  Select,
  TextField,
  Typography,
} from "@mui/material";
import dayjs from "dayjs";
import { useState } from "react";
import { useMutation } from "react-fetching-library";
import { Controller, SubmitHandler, useForm } from "react-hook-form";
import DatePicker from "../../components/datePicker";
import DateTimePicker from "../../components/dateTimePicker";
import FilesUpload from "../../components/filesUpload";
import { fetching } from "../../fetching";
import {
  AustraliaState,
  PropertyCategory,
  PropertyType,
} from "../../types/enum";
import { CreatePropertyRequest } from "../../types/types";

export default function CreatePropertyPage() {
  const {
    control,
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<CreatePropertyRequest>({
    defaultValues: {
      type: PropertyType.Unit,
      category: PropertyCategory.Rent,
      state: AustraliaState.WesternAustralia,
    },
  });
  const [uploadedFiles, setUploadedFiles] = useState<File[]>([]);
  const [isRenting, setIsRenting] = useState(true);
  const { mutate: upload } = useMutation(fetching.mutation.fileUpload);
  const { mutate } = useMutation(fetching.mutation.createProperty);
  const [submitting, setSubmitting] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [availableAt, setAvailableAt] = useState<Date>();
  const [openAt, setOpenAt] = useState<Date>();

  const handleAvailableAtChange = (selectedDateTime: dayjs.Dayjs | null) => {
    if (selectedDateTime) {
      const dateValue = selectedDateTime.toDate();
      setAvailableAt(dateValue);
    } else {
      setAvailableAt(undefined);
    }
  };

  const handleOpenAtChange = (selectedDateTime: dayjs.Dayjs | null) => {
    if (selectedDateTime) {
      const dateValue = selectedDateTime.toDate();
      setOpenAt(dateValue);
    } else {
      setOpenAt(undefined);
    }
  };

  const handleFilesUpload = (files: File[]) => {
    setUploadedFiles(files);
  };

  const onSubmit: SubmitHandler<CreatePropertyRequest> = async (input) => {
    setSubmitting(true);
    const property: CreatePropertyRequest = {
      ...input,
      bed_count: Number(input.bed_count),
      bath_count: Number(input.bath_count),
      car_count: Number(input.car_count),
      price: Number(input.price),
      has_aircon: typeof input.has_aircon === "string" ? false : true,
      is_furnished: typeof input.is_furnished === "string" ? false : true,
      is_pets_considered:
        typeof input.is_pets_considered === "string" ? false : true,
    };

    console.log(property.available_at);
    try {
      if (uploadedFiles.length > 0) {
        // The ! are changed to / in the backend (otherwise the path is removed by filename.base)
        const filePath = "album!album_art!";
        const res = await upload({
          file: uploadedFiles[0],
          public: true,
          filePath: filePath,
        });

        if (res.error) {
          setErrorMessage("Failed to upload image, try again.");
          setSubmitting(false);
        }
      }

      const resp = await mutate(property);
      if (resp.error) {
        setErrorMessage("Failed to create property, try again.");
        setSubmitting(false);
      }
    } catch (error) {
      setErrorMessage(
        typeof error === "string"
          ? error
          : "Something went wrong, please try again."
      );
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <Box
      sx={{
        maxWidth: "60em",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <form
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "2em",
        }}
        onSubmit={handleSubmit(onSubmit)}
        className="Form"
      >
        <div
          style={{
            display: "flex",
            gap: "1em",
          }}
        >
          <Typography fontWeight={"bold"}>Type</Typography>
          <FormControl fullWidth {...register("type")}>
            <Controller
              name="type"
              control={control}
              render={({ field }) => (
                <Select
                  {...field}
                  size="small"
                  onChange={(e) => {
                    field.onChange(e.target.value as PropertyType);
                  }}
                >
                  {Object.values(PropertyType).map((value) => (
                    <MenuItem key={"property-type" + value} value={value}>
                      {value}
                    </MenuItem>
                  ))}
                </Select>
              )}
            />
          </FormControl>
        </div>
        <div style={{ display: "flex", gap: "1em" }}>
          <Typography fontWeight={"bold"}>Category</Typography>
          <FormControl fullWidth {...register("category")}>
            <Controller
              name="category"
              control={control}
              render={({ field }) => (
                <Select
                  {...field}
                  size="small"
                  onChange={(e) => {
                    field.onChange(e.target.value as PropertyCategory);
                  }}
                >
                  {Object.values(PropertyCategory).map((value) => (
                    <MenuItem key={"property-category" + value} value={value}>
                      {value}
                    </MenuItem>
                  ))}
                </Select>
              )}
            />
          </FormControl>
        </div>
        <div>
          <div style={{ display: "flex", gap: "1em" }}>
            <Typography fontWeight={"bold"}>Street</Typography>
            <TextField
              size="small"
              fullWidth
              {...register("street", {
                required: {
                  value: true,
                  message: "Please enter street.",
                },
              })}
            />
          </div>
          {errors.street && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.street.message}
            </span>
          )}
        </div>
        <div>
          <div style={{ display: "flex", gap: "1em" }}>
            <Typography fontWeight={"bold"}>Suburb</Typography>
            <TextField
              size="small"
              fullWidth
              {...register("suburb", {
                required: {
                  value: true,
                  message: "Please enter suburb.",
                },
              })}
            />
          </div>
          {errors.suburb && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.suburb.message}
            </span>
          )}
        </div>
        <div>
          <div style={{ display: "flex", gap: "1em" }}>
            <Typography fontWeight={"bold"}>Postcode</Typography>
            <TextField
              size="small"
              fullWidth
              {...register("postcode", {
                required: {
                  value: true,
                  message: "Please enter postcode.",
                },
              })}
            />
          </div>
          {errors.postcode && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.postcode.message}
            </span>
          )}
        </div>

        <div style={{ display: "flex", gap: "1em" }}>
          <Typography fontWeight={"bold"}>State</Typography>
          <FormControl fullWidth {...register("state")}>
            <Controller
              name="state"
              control={control}
              render={({ field }) => (
                <Select
                  {...field}
                  size="small"
                  onChange={(e) => {
                    field.onChange(e.target.value as AustraliaState);
                  }}
                >
                  {Object.values(AustraliaState).map((value) => (
                    <MenuItem key={"state" + value} value={value}>
                      {value}
                    </MenuItem>
                  ))}
                </Select>
              )}
            />
          </FormControl>
        </div>
        <div>
          <div style={{ display: "flex", gap: "1em" }}>
            <Typography fontWeight={"bold"}>Bed count</Typography>
            <TextField
              size="small"
              type="number"
              fullWidth
              {...register("bed_count", {
                required: {
                  value: true,
                  message: "Please enter bed count.",
                },
              })}
            />
          </div>
          {errors.bed_count && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.bed_count.message}
            </span>
          )}
        </div>
        <div>
          <div style={{ display: "flex", gap: "1em" }}>
            <Typography fontWeight={"bold"}>Bath count</Typography>
            <TextField
              size="small"
              type="number"
              fullWidth
              {...register("bath_count", {
                required: {
                  value: true,
                  message: "Please enter bath count.",
                },
              })}
            />
          </div>
          {errors.bath_count && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.bath_count.message}
            </span>
          )}
        </div>
        <div style={{ display: "flex", gap: "1em" }}>
          <Typography fontWeight={"bold"}>Car count</Typography>
          <TextField
            size="small"
            type="number"
            fullWidth
            {...register("car_count", {
              required: {
                value: true,
                message: "Please enter car count.",
              },
            })}
          />
        </div>
        <div style={{ display: "flex", gap: "1em" }}>
          <Typography fontWeight={"bold"}>Aircon?</Typography>
          <FormControl fullWidth>
            <Controller
              name="has_aircon"
              control={control}
              defaultValue={true}
              render={({ field }) => (
                <RadioGroup {...field}>
                  <FormControlLabel
                    value={true}
                    control={<Radio />}
                    label="Yes"
                  />
                  <FormControlLabel
                    value={false}
                    control={<Radio />}
                    label="No"
                  />
                </RadioGroup>
              )}
            />
          </FormControl>
        </div>
        <div style={{ display: "flex", gap: "1em" }}>
          <Typography fontWeight={"bold"}>Furnished?</Typography>
          <FormControl fullWidth>
            <Controller
              name="is_furnished"
              control={control}
              defaultValue={true}
              render={({ field }) => (
                <RadioGroup {...field}>
                  <FormControlLabel
                    value={true}
                    control={<Radio />}
                    label="Yes"
                  />
                  <FormControlLabel
                    value={false}
                    control={<Radio />}
                    label="No"
                  />
                </RadioGroup>
              )}
            />
          </FormControl>
        </div>
        <div style={{ display: "flex", gap: "1em" }}>
          <Typography fontWeight={"bold"}> Pets?</Typography>
          <FormControl fullWidth>
            <Controller
              name="is_pets_considered"
              control={control}
              defaultValue={true}
              render={({ field }) => (
                <RadioGroup {...field}>
                  <FormControlLabel
                    value={true}
                    control={<Radio />}
                    label="Yes"
                  />
                  <FormControlLabel
                    value={false}
                    control={<Radio />}
                    label="No"
                  />
                </RadioGroup>
              )}
            />
          </FormControl>
        </div>
        <div style={{ display: "flex", gap: "1em" }}>
          <Typography fontWeight={"bold"}> Available at</Typography>
          <DatePicker handleChange={handleAvailableAtChange} />
        </div>
        <div style={{ display: "flex", gap: "1em" }}>
          <Typography fontWeight={"bold"}> Open at</Typography>
          <DateTimePicker handleChange={handleOpenAtChange} />
        </div>
        <div>
          <div style={{ display: "flex", gap: "1em" }}>
            <Typography fontWeight={"bold"}>
              {isRenting ? "Price(Weekly)" : "Price"}
            </Typography>
            <TextField
              size="small"
              fullWidth
              type="number"
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">$</InputAdornment>
                ),
              }}
              {...register("price", {
                required: {
                  value: true,
                  message: "Please enter price.",
                },
              })}
            />
          </div>
          {errors.price && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.price.message}
            </span>
          )}
        </div>
        <div style={{ display: "flex", gap: "1em" }}>
          <Typography fontWeight={"bold"}>Upload images</Typography>
          <FilesUpload onFilesUpload={handleFilesUpload} />
        </div>
        <Button variant="contained" type="submit" disabled={submitting}>
          Submit
        </Button>
        <Button variant="contained" disabled={submitting}>
          Cancel
        </Button>
      </form>
    </Box>
  );
}
