# Cloud Private Management System

## Description

This is a cloud private management system developed in Golang specifically designed for managing your exclusive cloud infrastructure. The system is tailored to work only within your own cloud environment and does not support management of cloud resources for other users or cloud providers. It provides a unified set of functionalities for interacting with your private cloud infrastructure through an API, application, or command-line interface (CLI). Whether you prefer programmatic control, a user-friendly graphical interface, or a command-line tool, this system has you covered.

## Features

- **API**: The system exposes a comprehensive API that enables programmatic access to your private cloud resources. You can use the API to automate tasks, integrate with other systems, or develop custom applications that interact with your cloud infrastructure.

- **Application**: The system includes a user-friendly application with a graphical interface specifically designed to manage your private cloud resources. The application provides a rich set of features, including a dashboard for monitoring resource usage, a virtual machine manager, a storage manager, and a network configuration tool. This application allows you to efficiently and intuitively manage your cloud infrastructure.

- **Command-Line Interface (CLI)**: The system offers an interactive CLI, providing a menu-driven interface within the terminal. The CLI allows you to perform all management tasks directly from the command line, providing a convenient way to manage your cloud infrastructure. The interactive menu guides you through the available options and commands, making it easy to perform various operations efficiently.

## Installation

### Prerequisites

Before installing the Cloud Private Management System, ensure that you have the following prerequisites:

- Golang (version X.X.X)
- [Cloud Provider SDK](link-to-sdk) (version X.X.X)

### Instructions

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/vinybergamo/cloud-manager.git
   ```

2. Change into the project directory:

   ```bash
   cd cloud-manager
   ```

3. Build the project:
   ```bash
   go build
   ```

## Usage

The Cloud Private Management System is designed to manage your exclusive cloud infrastructure. It supports different modes of operation. Choose the mode that suits your needs by specifying the appropriate command-line argument:

- To run in API mode:

  ```bash
  ./app --api
  ```

- To run in application mode:

  ```bash
  ./app --app
  ```

- To run in CLI mode:
  ```bash
  ./app
  ```

The CLI mode provides an interactive terminal interface where you can navigate through the available options and commands to manage your cloud infrastructure.

Please note that this system is specifically tailored to work within your own cloud environment and is not intended for managing cloud resources for other users or cloud providers.

## License

This project is not currently open-source and does not have a public license.
