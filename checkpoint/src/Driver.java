// this is the driver for the big data processing toolbox

import java.util.Scanner;

public class Driver {

    public static void main(String[] args) {
        System.out.println("Welcome to Big Data Processing Application\n" +
        "Please tpye the number that corresponding to which application you would like to run:\n" +
        "1. Apache Hadoop\n" +
        "2. Apache Spark\n" +
        "3. Jupyter Notebook\n" +
        "4. SonarQube and SonarScanner\n\n"+
        "(You can type anything other than 1-4 to quit the application)\n" +
        "Type the number here >");

        boolean toQuit = false;

        Scanner scanner = new Scanner(System.in);
        while(scanner.hasNextInt()) {
            int num = scanner.nextInt();
            switch (num) {
                case 1:
                    System.out.println("The URL for Apache Hadoop is ");
                    break;
                case 2:
                    System.out.println("The URL for Apache Spark is ");
                    break;
                case 3:
                    System.out.println("The URL for Jupyter Notebook is ");
                    break;
                case 4:
                    System.out.println("The URL for SonarQube and SonarScanner is ");
                    break;
                default:
                    System.out.println("The input number is out of bound.");
                    toQuit = true;
                    break;
            }
            if (toQuit) {
                break;
            }
        }
        scanner.close();
    }
}
