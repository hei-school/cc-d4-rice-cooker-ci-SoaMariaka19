module RiceCooker
    STATE_OFF = "Off"
    STATE_COOKING = "Cooking in progress"
    STATE_KEEP_WARM = "Keep warm mode"
  
    @rice_cooker_state = STATE_OFF if @rice_cooker_state.nil?
    @rice_and_water_added = false
  
    def self.display_state
      puts "\nCurrent state of the rice cooker: #{@rice_cooker_state}"
    end
  
    def self.plug_in
      if @rice_cooker_state == STATE_OFF && @rice_and_water_added
        @rice_cooker_state = STATE_COOKING
        puts "The rice cooker is plugged in, and cooking begins."
      elsif !@rice_and_water_added
        puts "Error: Add rice and water before plugging in and starting cooking."
      elsif @rice_cooker_state == STATE_COOKING || @rice_cooker_state == STATE_KEEP_WARM
        puts "Error: The rice cooker is already cooking or in keep warm mode."
      end
    end
  
    def self.finish_cooking
      if @rice_cooker_state == STATE_COOKING
        @rice_cooker_state = STATE_KEEP_WARM
        puts "Cooking is finished. The rice cooker is in keep warm mode."
      else
        puts "Error: No cooking in progress."
      end
    end
  
    def self.quit_program
      puts "\n======================================================================="
      puts "Goodbye! Thanks for using the rice cooker program.\n"
      exit
    end
  
    def self.set_rice_and_water_added(value)
      @rice_and_water_added = value
    end
  
    if __FILE__ == $PROGRAM_NAME
      puts "\nWelcome to the rice cooker program."
  
      loop do
        puts "\n=======================================================================\n        Menu:\n"
        puts "1. Add rice and water"
        puts "2. Plug in the rice cooker"
        puts "3. Cook rice"
        puts "4. Keep warm"
        puts "5. Rice cooker state"
        puts "6. End of cooking notification"
        puts "7. Quit the program"
  
        print "\n=======================================================================\nEnter your choice number: "
        choice = gets.chomp
  
        unless choice.match?(/^\d+$/)
          puts "Error: Please enter a number."
          next
        end
  
        choice_number = choice.to_i
  
        case choice_number
        when 1
          puts "You added rice and water."
          @rice_and_water_added = true
        when 2
          plug_in
        when 3
          if @rice_cooker_state == STATE_OFF
            puts "Error: Add rice and water before starting cooking."
          else
            puts "Cooking rice is in progress."
          end
        when 4
          if @rice_cooker_state == STATE_COOKING
            finish_cooking
          else
            puts "Error: No cooking in progress."
          end
        when 5
          display_state
        when 6
          if @rice_cooker_state == STATE_KEEP_WARM
            puts "Cooking is finished. The rice cooker is in keep warm mode."
          else
            puts "Error: No finished cooking."
          end
        when 7
          quit_program
        else
          puts "Error: Invalid choice. Please enter a valid number."
        end
      end
    end
  end
  