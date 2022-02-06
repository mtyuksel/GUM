# Git User Manager (GUM)

It is the project that enables to switch between Git users via terminal. It just uses Git commands and does not directly access Git files (like config file) for any operation.

## Installation

- Clone this repo to your PC:

```git
git clone https://github.com/mtyuksel/gum.git
```

- Open your zshrc/bashrc file with any editor:

```
code ~/.zshrc
```

- Add following with the correct directory at the and of the zshrc/bashrc file:
```
export PATH=$PATH:<THE_DIRECTORY_YOU_CLONED_GUM_INTO>/gum
```

- And now, you can reach GUM in any directory. Check with following (you will see default profile names):

![image](https://user-images.githubusercontent.com/60136172/152702436-8fb52777-4f0d-4130-bba5-0fcbab89dc89.png)

## Get Started

- Add your profile info to config.yaml

![image](https://user-images.githubusercontent.com/60136172/152701323-63fdf2a0-0d91-448f-89d3-8684944c7629.png)

- You can set your user profiles with following command:
```
gum set <profile-name>
```
![image](https://user-images.githubusercontent.com/60136172/152702464-f55a8cbc-4cbb-46df-a872-133a9aa6f94c.png)

- You can check your user profile names with following command:
```
gum show
```
![image](https://user-images.githubusercontent.com/60136172/152702440-c3f1910c-a7c7-4e5f-9b8c-d433afaa8817.png)

