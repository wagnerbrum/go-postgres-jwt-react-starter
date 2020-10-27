import React, { useState } from "react";
import { Link } from "react-router-dom";
import { useForm } from "react-hook-form";

import api from "../../services/api";

import "./index.css";

const Register = ({ history }) => {
    const [state, setState] = useState({
        email: "",
        password: "",
        name: "",
        isSubmitting: false,
        message: "",
        apiErrors: null,
    });

    const { register, handleSubmit, errors } = useForm();

    const { message, isSubmitting, apiErrors } = state;

    const onSubmit = async (data) => {
        setState({ ...state, isSubmitting: true, message: "" });

        const { input_email, input_password, input_name } = data;

        try {
            await api.post("/register", {
                email: input_email,
                password: input_password,
                name: input_name,
            });

            history.push("/login");
        } catch (e) {
            if (e.response && e.response.data) {
                const { msg, errors: apiErrors } = e.response.data;

                return setState({
                    ...state,
                    message: msg,
                    apiErrors,
                    isSubmitting: false,
                });
            }

            setState({ ...state, message: e.toString(), isSubmitting: false });
        }
    };

    return (
        <div className="container container-register">
            <div className="card text-center col-8 px-0">
                <div className="card-header">
                    <h2>Register</h2>
                </div>
                <div className="card-body">
                    <form onSubmit={handleSubmit(onSubmit)}>
                        <div className="row">
                            {message || apiErrors ? (
                                <div className="alert alert-warning col-8 mx-auto mt-2">
                                    {apiErrors ? (
                                        <div>
                                            {apiErrors.map((error, id) => {
                                                return (
                                                    <span key={id}>
                                                        {error}
                                                        <br />
                                                    </span>
                                                );
                                            })}
                                        </div>
                                    ) : (
                                        message
                                    )}
                                </div>
                            ) : (
                                <div className="alert alert-primary col-8 mx-auto mt-2">
                                    Enter with your informations
                                </div>
                            )}
                        </div>

                        <div className="row mt-2">
                            <div className="form-group col-8 mx-auto text-left">
                                <input
                                    type="text"
                                    className="form-control"
                                    placeholder="Email"
                                    name="input_email"
                                    ref={register({
                                        required: true,
                                    })}
                                />
                                <small className="text-danger ">
                                    {errors.input_email &&
                                        errors.input_email.type ===
                                            "required" &&
                                        "This input is required"}
                                </small>
                            </div>

                            <div className="form-group col-8 mx-auto text-left">
                                <input
                                    type="password"
                                    className="form-control"
                                    placeholder="Password"
                                    name="input_password"
                                    ref={register({
                                        required: true,
                                        minLength: 4,
                                    })}
                                />
                                <small className="text-danger ">
                                    {errors.input_password &&
                                        errors.input_password.type ===
                                            "required" &&
                                        "This input is required"}

                                    {errors.input_password &&
                                        errors.input_password.type ===
                                            "minLength" &&
                                        "This input must contain at least 4 characters"}
                                </small>
                            </div>

                            <div className="form-group col-8 mx-auto text-left">
                                <input
                                    type="text"
                                    className="form-control"
                                    placeholder="Name"
                                    name="input_name"
                                    ref={register({ required: true })}
                                />
                                <small className="text-danger ">
                                    {errors.input_name &&
                                        errors.input_name.type === "required" &&
                                        "This input is required"}
                                </small>
                            </div>
                        </div>

                        <button
                            type="submit"
                            className="btn btn-primary mx-auto my-2"
                            disabled={isSubmitting}
                        >
                            {isSubmitting ? "Loading ..." : "Register"}
                        </button>

                        <Link to="/" className="mt-2">
                            Back to Login
                        </Link>
                    </form>
                </div>
            </div>
        </div>
    );
};

export default Register;
